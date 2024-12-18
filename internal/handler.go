package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/miekg/dns"
	"go.uber.org/zap"
)

type DNSStatus string

const (
	DNSStatusUp      DNSStatus = "UP"
	DNSStatusDown    DNSStatus = "DOWN"
	DNSStatusUnknown DNSStatus = "UNKNOWN"
)

type URLScanResponse struct {
	Results []URLScanResult `json:"results"`
}

type URLScanResult struct {
	URL         string    `json:"url"`
	DNSStatus   DNSStatus `json:"dns_status"`
	Categories  []string  `json:"categories,omitempty"`
	IsMalicious bool      `json:"is_malicious,omitempty"`
}

type Handler struct {
	logger *zap.Logger
}

func NewHandler(logger *zap.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

// getDomainCategories returns categories and malicious status for a domain
// will be replaced with proper service call in the future
func (h *Handler) getDomainCategories(domain string) ([]string, bool) {
	// Clean up the domain (remove scheme, path, etc.)
	cleanDomain := domain
	if strings.Contains(domain, "://") {
		if parsedURL, err := url.Parse(domain); err == nil {
			cleanDomain = parsedURL.Host
		}
	}

	// Remove www. prefix if present
	cleanDomain = strings.TrimPrefix(cleanDomain, "www.")

	// First check if domain is malicious
	isMalicious := MaliciousDomains[cleanDomain]

	// Get categories
	categories := DomainCategories[cleanDomain]

	// If no direct match, try to match parent domain
	if len(categories) == 0 {
		parts := strings.Split(cleanDomain, ".")
		if len(parts) > 2 {
			parentDomain := strings.Join(parts[len(parts)-2:], ".")
			categories = DomainCategories[parentDomain]
		}
	}

	return categories, isMalicious
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var responseList URLScanResponse
	// Validation
	if r.Method != http.MethodGet {
		h.handleError(w, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlParam := r.URL.Query().Get("url")
	if urlParam == "" {
		h.handleError(w, fmt.Errorf("URL is required"), http.StatusBadRequest)
		return
	}
	dnsStatusFlag := r.URL.Query().Get("dns_status")
	if dnsStatusFlag != "" && dnsStatusFlag != "1" {
		h.handleError(w, fmt.Errorf("invalid dns_status flag. Only '1' is accepted"), http.StatusBadRequest)
		return
	}
	categoriesFlag := r.URL.Query().Get("categories")
	if categoriesFlag != "" && categoriesFlag != "1" {
		h.handleError(w, fmt.Errorf("invalid categories flag. Only '1' is accepted"), http.StatusBadRequest)
		return
	}

	// Split URLs
	urls := strings.Split(urlParam, ",")
	for i, u := range urls {
		urls[i] = strings.TrimSpace(u)
	}

	// Perform the scan
	for _, url := range urls {
		response, err := h.isUP(url)
		if err != nil {
			h.handleError(w, err, http.StatusInternalServerError)
			return
		}

		// Add categories and malicious status if categories flag is set
		if categoriesFlag == "1" {
			categories, isMalicious := h.getDomainCategories(url)
			response.Categories = categories
			response.IsMalicious = isMalicious
		}

		responseList.Results = append(responseList.Results, *response)
		h.logger.Info("URL scanned",
			zap.String("url", response.URL),
			zap.String("status", string(response.DNSStatus)),
			zap.String("categories", strings.Join(response.Categories, ",")),
			zap.Bool("is_malicious", response.IsMalicious),
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseList)
}

// isUP performs the actual scanning logic
func (h *Handler) isUP(inputUrl string) (*URLScanResult, error) {
	// Extract the hostname from the input
	var domain string
	if parsedURL, err := url.Parse(inputUrl); err == nil && parsedURL.Host != "" {
		domain = parsedURL.Host
	} else {
		domain = inputUrl
	}

	// Ensure the domain is fully qualified
	if !strings.HasSuffix(domain, ".") {
		domain = dns.Fqdn(domain)
	}
	var dnsStatus DNSStatus = DNSStatusUnknown
	c := new(dns.Client)
	c.Timeout = 5 * time.Second

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeA)

	r, _, err := c.Exchange(m, "127.0.0.1:53")
	if err != nil {
		h.logger.Error("DNS exchange error", zap.Error(err))
		return &URLScanResult{
			URL:        domain,
			DNSStatus:  DNSStatusUnknown,
			Categories: []string{},
		}, nil
	}

	switch r.Rcode {
	case dns.RcodeSuccess:
		if len(r.Answer) > 0 {
			dnsStatus = DNSStatusUp
		} else {
			h.logger.Info("DNS query successful but no A records found", zap.String("url", domain))
			dnsStatus = DNSStatusDown
		}
	case dns.RcodeNameError, // NXDOMAIN
		dns.RcodeServerFailure, // SERVFAIL
		dns.RcodeRefused,       // REFUSED
		dns.RcodeNotAuth,       // NOTAUTH
		dns.RcodeNotZone:       // NOTZONE
		h.logger.Info("DNS query failed", zap.String("dns code", strconv.Itoa(r.Rcode)))
		dnsStatus = DNSStatusDown
	default:
		dnsStatus = DNSStatusUnknown
	}

	response := &URLScanResult{
		URL:        domain,
		DNSStatus:  dnsStatus,
		Categories: []string{}, // Will be populated later if categories flag is set
	}
	h.logger.Info("DNS scan result",
		zap.String("input", inputUrl),
		zap.String("domain", domain),
		zap.String("status", string(dnsStatus)))

	return response, nil
}

// handleError handles error responses
func (h *Handler) handleError(w http.ResponseWriter, err error, statusCode int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

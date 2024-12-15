package internal

import (
	"context"
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

type URLScanRequest struct {
	URL string `json:"url"`
}

type URLScanResponse struct {
	URL        string    `json:"url"`
	DNSStatus  DNSStatus `json:"dns_status"`
	Categories []string  `json:"categories,omitempty"`
}

type Handler struct {
	logger *zap.Logger
}

func NewHandler(logger *zap.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Validation
	if r.Method != http.MethodGet {
		h.handleError(w, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	url := r.URL.Query().Get("url")
	if url == "" {
		h.handleError(w, fmt.Errorf("URL is required"), http.StatusBadRequest)
		return
	}

	// Perform the scan
	response, err := h.scanURL(r.Context(), url)
	if err != nil {
		h.handleError(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	h.logger.Info("URL scanned",
		zap.String("url", response.URL),
		zap.String("status", string(response.DNSStatus)),
	)
}

// scanURL performs the actual scanning logic
func (h *Handler) scanURL(ctx context.Context, inputUrl string) (*URLScanResponse, error) {
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
		return &URLScanResponse{
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

	response := &URLScanResponse{
		URL:        domain,
		DNSStatus:  dnsStatus,
		Categories: []string{}, // Empty for now, to be implemented later
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

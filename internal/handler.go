package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/miekg/dns"
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

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Validattion
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
}

// scanURL performs the actual scanning logic
func (h *Handler) scanURL(ctx context.Context, url string) (*URLScanResponse, error) {
	var domain DNSStatus
	c := new(dns.Client)
	c.Timeout = 5 * time.Second

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(url), dns.TypeA)

	r, _, err := c.Exchange(m, "127.0.0.1:53")
	if err != nil {
		domain = DNSStatusUnknown
	}

	switch r.Rcode {
	case dns.RcodeSuccess:
		if len(r.Answer) > 0 {
			domain = DNSStatusUp
		}
		domain = DNSStatusDown // No valid answers, consider it down
	case dns.RcodeNameError: // NXDOMAIN
	case dns.RcodeServerFailure: // SERVFAIL
	case dns.RcodeRefused: // REFUSED
	case dns.RcodeNotAuth: // NOTAUTH
		domain = DNSStatusDown
	case dns.RcodeFormatError: // FORMERR
	case dns.RcodeNotImplemented: // NOTIMP
	case dns.RcodeNotZone: // NOTZONE
		domain = DNSStatusDown
	default:
		domain = DNSStatusUnknown
	}

	response := &URLScanResponse{
		URL:        url,
		DNSStatus:  domain,
		Categories: []string{}, // Empty for now, to be implemented later
	}

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

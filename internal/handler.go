package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
	//w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// scanURL performs the actual scanning logic
func (h *Handler) scanURL(ctx context.Context, url string) (*URLScanResponse, error) {
	response := &URLScanResponse{
		URL:        url,
		DNSStatus:  DNSStatusUnknown,
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

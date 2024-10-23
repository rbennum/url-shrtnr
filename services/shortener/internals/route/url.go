package route

import (
	"context"
	"encoding/json"
	"net/http"

	resp "github.com/rbennum/url-shrtnr/internals/models"
	"github.com/rbennum/url-shrtnr/internals/service"
	"github.com/rs/zerolog/log"
)

type UrlHandler struct {
	ctx        context.Context
	urlService service.UrlService
}

func NewUrlHandler(ctx context.Context, urlService service.UrlService) *UrlHandler {
	return &UrlHandler{ctx: ctx, urlService: urlService}
}

// Implementing ServeHTTP net/http/Handler
func (h *UrlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handleShortenUrl(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UrlHandler) handleShortenUrl(w http.ResponseWriter, r *http.Request) {
	var request struct {
		OriginalUrl string `json:"original_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	link, err := h.urlService.CreateUrl(h.ctx, request.OriginalUrl)
	if err != nil {
		log.Error().Err(err).Msg("Failed to shorten URL")
		response := resp.Response{
			Success: false,
			Message: "Failed to shorten URL",
			Error: &resp.ErrorInfo{
				Code:    "CREATE_URL_ERROR",
				Message: "Error creating new URL. Please try again.",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	response := resp.Response{
		Success: true,
		Message: "URL successfully shortened.",
		Data: &resp.Data{
			ShortURL:    link.Tag,
			OriginalURL: link.URL,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

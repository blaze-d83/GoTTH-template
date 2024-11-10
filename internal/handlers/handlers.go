package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/blaze-d83/go-GoTTH/internal/repository"
	"github.com/blaze-d83/go-GoTTH/internal/templates"
	"github.com/blaze-d83/go-GoTTH/pkg/logger"
)

type Handler struct {
	db      *sql.DB
	logger  logger.Logger
	queries *repository.Queries
}

func NewHandler(db *sql.DB, l logger.Logger) *Handler {
	return &Handler{
		db:      db,
		logger:  l,
		queries: repository.New(db),
	}
}

func (h Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	homePage := templates.BaseTemplate()
	c := r.Context()
	if err := homePage.Render(c, w); err != nil {
		h.logger.LogError(c, err, r.Method, r.URL.Path, r.RequestURI)
		http.Error(w, `{"error": "Failed to render homepage"}`, http.StatusInternalServerError)
	}
}

func (h *Handler) GetCounter(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	c := r.Context()
	counter, err := h.queries.GetCounter(c)
	duration := time.Since(startTime)

	h.logger.LogRequests(c, r.Method, r.URL.Path, r.RemoteAddr, r.RequestURI)

	if err != nil {
		h.logger.LogError(c, err, r.Method, r.URL.Path, r.RequestURI)
		http.Error(w, `{"error": "Failed to get counter"}`, http.StatusInternalServerError)
	} else {
		h.logger.LogEvent(c, "Fetched counter values successfully")
		h.logger.LogResponses(c, r.Response.StatusCode, duration, r.Method, r.URL.Path, r.RequestURI)
	}

	response := map[string]int64{"count": counter}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) IncrementCounter(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	h.logger.LogRequests(c, r.Method, r.URL.Path, r.RemoteAddr, r.RequestURI)
	if err := h.queries.IncrementCounter(r.Context()); err != nil {
		h.logger.LogError(c, err, r.Method, r.URL.Path, r.RequestURI)
		http.Error(w, `{"error": "Failed to increment counter"}`, http.StatusInternalServerError)
		return
	} else {
		h.logger.LogEvent(c, "Counter incremented successfully")
	}

}

func (h *Handler) DecrementCounter(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	h.logger.LogRequests(c, r.Method, r.URL.Path, r.RemoteAddr, r.RequestURI)
	if err := h.queries.DecrementCounter(c); err != nil {
		h.logger.LogError(c, err, r.Method, r.URL.Path, r.RequestURI)
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	response := map[string]string{"status": "decrement"}
	json.NewEncoder(w).Encode(response)
}

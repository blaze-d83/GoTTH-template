package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/blaze-d83/go-GoTTH/internal/repository"
	"github.com/blaze-d83/go-GoTTH/internal/templates"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	db      *sql.DB
	logger  *logrus.Logger
	queries *repository.Queries
}

func NewHandler(db *sql.DB, logger *logrus.Logger) *Handler {
	return &Handler{
		db:      db,
		queries: repository.New(db),
	}
}

func (h Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	homePage := templates.BaseTemplate()
	c := r.Context()
	if err := homePage.Render(c, w); err != nil {
		h.logger.WithError(err).Error("Failed to render homepage")
		http.Error(w, `{"error": "Failed to render homepage"}`, http.StatusInternalServerError)
		return
	}

}

func (h *Handler) GetCounter(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	counter, err := h.queries.GetCounter(c)
	if err != nil {
		http.Error(w, `{"error": "Failed to get counter"}`, http.StatusInternalServerError)
		return
	}
	response := map[string]int64{"count": counter}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) IncrementCounter(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	if err := h.queries.IncrementCounter(c); err != nil {
		http.Error(w, `{"error": "Failed to increment counter"}`, http.StatusInternalServerError)
		return
	}

}

func (h *Handler) DecrementCounter(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	// Attempt to decrement the counter
	if err := h.queries.DecrementCounter(c); err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	response := map[string]string{"status": "decrement"}
	json.NewEncoder(w).Encode(response)
}

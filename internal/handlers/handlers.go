package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/blaze-d83/go-GoTTH/internal/service"
	"github.com/blaze-d83/go-GoTTH/internal/templates"
)

type Handler struct {
	service *service.CounterService
}

func NewHandler(service *service.CounterService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h Handler) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	homePage := templates.BaseTemplate()
	if err := homePage.Render(context.Background(), w); err != nil {
		http.Error(w, `{"error": "Failed to render homepage"}`, http.StatusInternalServerError)
	}
}

func (h *Handler) GetCounter(w http.ResponseWriter, r *http.Request) {
	count, err := h.service.GetCounter(context.Background())
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
	}
	response := map[string]int64{"count": count}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) UpdateCounter(w http.ResponseWriter, r *http.Request) {
	if err := h.service.IncrementCounter(context.Background()); err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	response := map[string]string{"status": "increment"}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DecrementCounter(w http.ResponseWriter, r *http.Request) {
	if err := h.service.DecrementCounter(context.Background()); err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	response := map[string]string{"status": "decrement"}
	json.NewEncoder(w).Encode(response)
}

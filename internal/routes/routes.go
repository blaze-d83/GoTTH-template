package routes

import (
	"net/http"

	"github.com/blaze-d83/go-GoTTH/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux, h *handlers.Handler) {
	mux.HandleFunc("/home", h.RenderHomePage)
	mux.HandleFunc("/counter", h.GetCounter)
	mux.HandleFunc("/update_counter", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			h.UpdateCounter(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/decrement_counter", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			h.DecrementCounter(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

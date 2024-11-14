package internal

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func RegisterRoutes(mux *http.ServeMux, handler *Handler, logger *logrus.Logger ) http.Handler  {
	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static", fs))
	mux.HandleFunc("/", handler.HomePage)
	mux.HandleFunc("/counter", handler.GetCounter)
	mux.HandleFunc("/increment", handler.IncrementCounter)
	mux.HandleFunc("/decrement", handler.DecrementCounter)

	loggedMux := LoggingMiddleware(logger, mux)

	return loggedMux
}

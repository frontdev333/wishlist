package home

import (
	"io"
	"net/http"
)

type Handler struct {
}

func (h *Handler) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Welcome to homePage")
	}
}

func (h *Handler) Configure(r *http.ServeMux) {
	r.HandleFunc("/", h.Index())
}

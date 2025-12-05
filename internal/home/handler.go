package home

import (
	"io"
	"net/http"
)

type HomeHandler struct {
}

func (h *HomeHandler) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Welcome to homePage")
	}
}

func (h *HomeHandler) Configure(r *http.ServeMux) {
	r.HandleFunc("/", h.Index())
}

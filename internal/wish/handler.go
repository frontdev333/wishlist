package wish

import "net/http"

type Handler struct {
	repository *Repository
}

func (h *Handler) GetAll(userId uint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) Get(wishId uint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

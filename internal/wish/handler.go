package wish

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
		var payload StorePayload

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			fmt.Println(err)
			return
		}

		wish, err := h.repository.Store(&payload)
		if err != nil {
			fmt.Println(err)
			return
		}
		_ = wish
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

package internal

import (
	"frontdev333/wishlist/internal/home"
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	configureRouter(router)
	return router
}

func configureRouter(r *http.ServeMux) {
	homeHandler := home.HomeHandler{}
	homeHandler.Configure(r)
}

package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/terjeee/go_htmx/internal/middleware"
)

// DETTE ER EN FUNKSJOIN
func Handler(router *chi.Mux) {
	router.Use(chimiddle.StripSlashes)

	router.Route("/account", func(router chi.Router) {
		// Nighclub bouncer,
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}

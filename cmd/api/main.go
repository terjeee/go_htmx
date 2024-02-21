package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/terjeee/go_htmx/internal/handlers"
)

func main() {
	log.SetReportCaller(true)
	var router *chi.Mux = chi.NewRouter()
	handlers.Handler(router)

	fmt.Println("Starting GO_API.service")

	error := http.ListenAndServe("localhost:8000", router)

	if error != nil {
		log.Error(error)
	}
}

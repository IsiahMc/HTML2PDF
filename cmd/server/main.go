package main

import (
	"net/http"
	"time"

	"github.com/IsiahMc/HTML2PDF/internal/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(10, 1*time.Minute))
	// GET
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to HTML 2 PDF"))
	})
	// POST
	r.Post("/convert", api.ConvertHandler)
	http.ListenAndServe(":3000", r)

}

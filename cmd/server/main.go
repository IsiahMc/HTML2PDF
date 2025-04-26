package main

import (
	"net/http"

	"github.com/IsiahMc/HTML2PDF/internal/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	// GET
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	// POST
	r.Post("/convert", api.ConvertHandler)
	http.ListenAndServe(":3000", r)

}

package main

import (
	"github.com/go-chi/chi/v5"
	repo "github.com/jack/ecom/internal/adapters/postgresql/sqlc"
	"github.com/jack/ecom/internal/products"
	"github.com/jackc/pgx/v5"
)

func routes(db *pgx.Conn) chi.Router {
	r := chi.NewRouter()

	productService := products.NewService(repo.New(db))
	productHandler := products.NewHandler(productService)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", productHandler.ListProducts)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.FindProductById)
	})

	return r
}

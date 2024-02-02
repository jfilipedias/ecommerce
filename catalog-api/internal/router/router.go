package router

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jfilipedias/ecommerce/catalog-api/internal/category"
	"github.com/jfilipedias/ecommerce/catalog-api/internal/product"
)

func NewRouter(db *sql.DB) *chi.Mux {
	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(*categoryRepository)
	categoryHandler := category.NewHandler(categoryService)

	productRepository := product.NewRepository(db)
	productService := product.NewService(*productRepository)
	productHandler := product.NewHandler(productService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/categories", categoryHandler.Create)
	r.Get("/categories", categoryHandler.GetAll)
	r.Get("/categories/{id}", categoryHandler.GetByID)

	r.Post("/products", productHandler.Create)
	r.Get("/products", productHandler.GetAll)
	r.Get("/products/{id}", productHandler.GetByID)
	r.Get("/products/categories/{categoryID}", productHandler.GetByCategoryID)

	return r
}

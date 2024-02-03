package router

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jfilipedias/ecommerce/api/internal/handler"
	"github.com/jfilipedias/ecommerce/api/internal/repository"
	"github.com/jfilipedias/ecommerce/api/internal/service"
)

func NewRouter(db *sql.DB) *chi.Mux {
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(*categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(*productRepository)
	productHandler := handler.NewProductHandler(productService)

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

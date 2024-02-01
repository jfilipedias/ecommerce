package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jfilipedias/ecommerce-fullcycle/api/internal/handler"
	"github.com/jfilipedias/ecommerce-fullcycle/api/internal/repository"
	"github.com/jfilipedias/ecommerce-fullcycle/api/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

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

	p := os.Getenv("API_PORT")
	fmt.Printf("Server is running on port %q", p)
	http.ListenAndServe(p, r)
}

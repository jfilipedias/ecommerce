package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jfilipedias/ecommerce-fullcycle/catalog-api/database"
	"github.com/jfilipedias/ecommerce-fullcycle/catalog-api/internal/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db, err := database.NewDatabase()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	port := os.Getenv("API_PORT")
	router := router.NewRouter(db)
	fmt.Printf("Server is running on port %q\n", port)
	http.ListenAndServe(port, router)
}

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jfilipedias/ecommerce/api/db"
	"github.com/jfilipedias/ecommerce/api/internal/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	db, err := db.NewPostgresDatabase()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	port := os.Getenv("API_PORT")
	router := router.NewRouter(db)
	fmt.Printf("Server is running on port %q\n", port)
	http.ListenAndServe(port, router)
}

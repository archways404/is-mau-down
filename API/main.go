package main

import (
	"api/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    port := os.Getenv("PORT")

    http.HandleFunc("/test", handlers.TestEndpoint)
    http.HandleFunc("/mau", handlers.CheckURLEndpoint)
    log.Printf("Server is running on http://localhost:%s", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

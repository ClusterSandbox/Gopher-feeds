package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	fmt.Println("Welcome to the web crawler")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Printf("Server starting on port:%v", portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running on PORT:", portString)

	log.Fatal(srv.ListenAndServe())
}

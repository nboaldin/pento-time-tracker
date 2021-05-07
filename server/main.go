package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/nboaldin/pento-time-tracker/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(header, methods, origins)(r)))
}

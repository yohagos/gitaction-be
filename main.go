package main

import (
	"log"
	"net/http"
	"yohagos/gitaction-be/router"

	"github.com/gorilla/handlers"
)

func main() {
	router := router.NewRouter()
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:4200"})
	err := http.ListenAndServe(":8888", handlers.CORS(credentials, methods, origins)(router))
	if err != nil {
		log.Panic(err)
	}
}

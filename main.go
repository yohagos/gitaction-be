package main

import (
	"log"
	"net/http"
	"yohagos/gitaction-be/router"
)

func main() {
	router := router.NewRouter()
	err := http.ListenAndServe(":8888", router)
	if err != nil {
		log.Panic(err)
	}
}
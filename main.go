package main

import (
	"log"
	"net/http"

	"meteo/helpers"
	"meteo/routes"
)

func main() {
	helpers.ChargerTemplates()

	mux := routes.Creer()

	log.Println("Site lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

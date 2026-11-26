package main

import (
	"fmt"
	"net/http"

	"meteo/helpers"
	"meteo/routes"
)

func main() {
	helpers.ChargerTemplates()
	routes.Creer()

	fmt.Println("Site lancé sur http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}

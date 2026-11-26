package main

import (
	"fmt"
	"net/http"

	"meteo/helpers"
	"meteo/routes"
)

func main() {
	helpers.ChargerTemplates()

	mux := routes.Creer()

	fmt.Println("Site lancé sur http://localhost:8080")
	if err != nil {
		fmt.Fatal(http.localhost:8080, err)
		return
	}
}


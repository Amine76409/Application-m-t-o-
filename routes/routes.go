package routes

import (
	"fmt"
	"net/http"

	"meteo/controllers"
)

func Creer() {
	fmt.Println("Création des routes...")

	http.HandleFunc("GET /{$}", controllers.Accueil)
	fmt.Println("Route ajoutée : accueil")

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	fmt.Println("Route ajoutée : fichiers static")

	http.HandleFunc("/", controllers.PageIntrouvable)
	fmt.Println("Route ajoutée : page 404")
}

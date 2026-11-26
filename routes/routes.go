package routes

import (
	"net/http"

	"meteo/controllers"
)

func Creer() {
	http.HandleFunc("GET /{$}", controllers.Accueil)
	http.HandleFunc("GET /ville", controllers.Villes)
	http.HandleFunc("GET /meteo", controllers.Meteo)

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", controllers.PageIntrouvable)
}

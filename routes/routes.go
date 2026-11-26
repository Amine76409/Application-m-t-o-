package routes

import (
	"net/http"

	"meteo/controllers"
)

func Creer() {
	http.HandleFunc("GET /{$}", controllers.Accueil)
	http.HandleFunc("GET /ville", controllers.ListeVilles)
	http.HandleFunc("GET /meteo", controllers.ResultatMeteo)

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", controllers.PageIntrouvable)
}

package controllers

import (
	"fmt"
	"net/http"

	"meteo/helpers"
)

func Accueil(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Quelqu'un visite la page d'accueil")
	helpers.Rendre(w, "index.html", 200, nil)
}

func PageIntrouvable(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Page introuvable :", r.URL.Path)

	data := map[string]string{"Message": "Cette page n'existe pas."}
	helpers.Rendre(w, "erreur.html", 404, data)
}

package controllers

import (
	"fmt"
	"net/http"

	"meteo/helpers"
	"meteo/services"
)

func Meteo(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	fmt.Println("Page météo, code :", code)

	if code == "" {
		data := map[string]string{"Message": "Cette commune n'existe pas."}
		helpers.Rendre(w, "erreur.html", 404, data)
		return
	}

	commune, trouve, err := services.ChercherCommuneParCode(code)
	if err != nil {
		data := map[string]string{"Message": "Service momentanément indisponible."}
		helpers.Rendre(w, "erreur.html", 503, data)
		return
	}

	if !trouve {
		data := map[string]string{"Message": "Cette commune n'existe pas."}
		helpers.Rendre(w, "erreur.html", 404, data)
		return
	}

	jours, err := services.RecupererMeteo(commune)
	if err != nil {
		data := map[string]string{"Message": "Service momentanément indisponible."}
		helpers.Rendre(w, "erreur.html", 503, data)
		return
	}

	data := map[string]any{"Commune": commune, "Jours": jours}
	helpers.Rendre(w, "meteo.html", 200, data)
}

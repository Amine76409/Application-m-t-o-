package services

import (
	"strings"

	"meteo/api"
	"meteo/models"
)

func NormaliserSaisie(saisie string) string {
	mots := strings.Fields(saisie)
	return strings.Join(mots, " ")
}

func ValiderSaisie(saisie string) string {
	taille := len([]rune(saisie))

	if taille == 0 {
		return "Veuillez saisir le nom d'une ville."
	}
	if taille < 2 || taille > 60 {
		return "Le nom doit contenir entre 2 et 60 caractères."
	}
	return ""
}

func RechercherCommunes(nom string) ([]models.Commune, error) {
	return api.ChercherCommunes(nom)
}

func ChercherCommuneParCode(code string) (models.Commune, bool, error) {
	return api.ChercherCommuneParCode(code)
}

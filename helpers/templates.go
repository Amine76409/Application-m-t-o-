package helpers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var Templates *template.Template

func ChargerTemplates() {
	fmt.Println("Chargement des templates...")

	var err error
	Templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println("Erreur :", err)
		panic(err)
	}

	fmt.Println("Templates chargés")
}

func Rendre(w http.ResponseWriter, page string, code int, data map[string]string) {
	fmt.Println("Affichage de la page :", page, "| code :", code)

	var buf bytes.Buffer

	err := Templates.ExecuteTemplate(&buf, page, data)
	if err != nil {
		fmt.Println("Erreur dans le template :", err)
		http.Error(w, "Une erreur est survenue.", 500)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	buf.WriteTo(w)

	fmt.Println("Page envoyée")
}

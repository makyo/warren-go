package handlers

import (
	"html/template"
	"net/http"
)

func Front(w http.ResponseWriter) {
	t := template.Must(template.ParseFiles(
		"templates/front.tmpl",
		"templates/base.tmpl"))
	t.ExecuteTemplate(w, "base", nil)
}

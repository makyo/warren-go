package handlers

import (
	"html/template"
	"net/http"
)

func Front(w http.ResponseWriter, t *template.Template) {
	t.ExecuteTemplate(w, "base", nil)
}

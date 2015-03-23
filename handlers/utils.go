package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
)

func init() {
	gob.Register(Flash{})
}

/* Flash handling ---------------------------------------------------------- */
type Flash struct {
	Class   string
	Message string
}

func NewFlash(message string, class ...string) Flash {
	if len(class) == 0 {
		class = []string{"notice"}
	}
	return Flash{
		Class:   class[0],
		Message: message,
	}
}

/* Template rendering ------------------------------------------------------ */
func (h *Handlers) render(w http.ResponseWriter, r *http.Request, tmpl string, args ...interface{}) {
	t := template.Must(template.ParseFiles(
		fmt.Sprintf("templates/%s", tmpl),
		"templates/base.tmpl"))
	flashes := h.session.Flashes()
	if len(flashes) > 0 {
		h.session.Save(r, w)
	}
	templateArg := struct {
		Flashes []interface{}
		User    User
	}{
		Flashes: flashes,
		User:    h.user,
	}
	t.ExecuteTemplate(w, "base", templateArg)
}

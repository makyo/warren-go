// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

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

// Creates a new flash item to be stored in the session, with an optional CSS
// class, which can be "success", "info", "warning", or "danger".
func NewFlash(message string, class ...string) Flash {
	if len(class) == 0 {
		class = []string{"info"}
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

// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package handlers

import (
	"log"
	"net/http"

	"github.com/makyo/warren-go/models"
)

func (h *Handlers) SessionMiddleware(w http.ResponseWriter, r *http.Request, l *log.Logger) {
	session, err := h.sessionStore.Get(r, "warren-session")
	if err != nil {
		log.Print(err)
		http.Error(w, "An error occurred", 500)
	}
	h.session = session
}

func (h *Handlers) AuthenticationMiddleware(w http.ResponseWriter, r *http.Request, l *log.Logger) {
	auth := h.session.Values["authenticated"]
	if auth == nil {
		auth = false
	}
	username := h.session.Values["username"]
	if username == nil {
		username = ""
	}
	model := models.User{}
	if username != "" {
		var err error
		model, err = models.GetUser(username.(string), h.db)
		if err != nil {
			l.Print(err.Error())
			http.Error(w, "Could not fetch user", http.StatusInternalServerError)
		}
	}
	h.user = User{
		IsAuthenticated: auth.(bool),
		Model:           model,
	}
}

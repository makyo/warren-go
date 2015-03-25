// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package handlers

import (
"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/makyo/warren-go/models"
)

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

func (h *Handlers) CSRFMiddleware(w http.ResponseWriter, r *http.Request, l *log.Logger) {
	token := h.session.Values["_csrf_token"]
	if token == nil {
		h.session.Values["_csrf_token"] = fmt.Sprintf("%d", rand.Int63())
		token = h.session.Values["_csrf_token"]
		h.session.Save(r, w)
	}
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Print(err.Error())
			http.Error(w, "Could not parse form", http.StatusInternalServerError)
			return
		}
		if token != r.FormValue("_csrf_token") {
			http.Error(w, "CSRF failure", http.StatusForbidden)
			return
		}
		h.session.Values["_csrf_token"] = nil
		h.session.Save(r, w)
	}
}

func (h *Handlers) SessionMiddleware(w http.ResponseWriter, r *http.Request, l *log.Logger) {
	session, err := h.sessionStore.Get(r, "warren-session")
	if err != nil {
		log.Print(err)
		http.Error(w, "An error occurred", 500)
	}
	h.session = session
}

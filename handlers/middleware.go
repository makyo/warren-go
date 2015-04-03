// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/martini-contrib/render"

	"github.com/warren-community/warren/models"
)

// Provide authentication middleware that fetches the user from the database if
// they are currently logged in.
func (h *Handlers) AuthenticationMiddleware(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
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
			l.Printf("Could not fetch user: %+v\n", err)
			h.InternalServerError(w, r, render)
		}
	}
	h.user = User{
		IsAuthenticated: auth.(bool),
		Model:           model,
	}
}

// Provide cross-site request forgery protection through middleware, ensuring
// that forms posted to the site contain the correct token.
func (h *Handlers) CSRFMiddleware(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	token := h.session.Values["_csrf_token"]
	if token == nil {
		h.session.Values["_csrf_token"] = fmt.Sprintf("%d", rand.Int63())
		token = h.session.Values["_csrf_token"]
		h.session.Save(r, w)
	}
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			l.Printf("Could not parse form: %+v\n", err)
			h.InternalServerError(w, r, render)
			return
		}
		if token != r.FormValue("_csrf_token") {
			h.BadRequest(w, r, render)
			return
		}
		h.session.Values["_csrf_token"] = nil
		h.session.Save(r, w)
	}
}

func (h *Handlers) SessionMiddleware(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	session, err := h.sessionStore.Get(r, "warren-session")
	if err != nil {
		log.Printf("Session error: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	h.session = session
}

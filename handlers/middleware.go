package handlers

import (
	"log"
	"net/http"
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
	h.user = User{
		IsAuthenticated: auth.(bool),
		Username:        username.(string),
	}
}

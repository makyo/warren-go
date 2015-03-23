package handlers

import (
	"github.com/gorilla/sessions"

	"log"
	"net/http"
)

type User struct {
	IsAuthenticated bool
	Username        string
}

type Handlers struct {
	sessionStore sessions.Store
	user         User
}

func New(store sessions.Store) Handlers {
	h := Handlers{
		sessionStore: store,
	}
	return h
}

func (h *Handlers) AuthenticationMiddleware(w http.ResponseWriter, r *http.Request, l *log.Logger) {
	session, err := h.sessionStore.Get(r, "user-management")
	if err != nil {
		log.Print(err)
		http.Error(w, "An error occurred", 500)
	}
	auth := session.Values["authenticated"]
	if auth == nil {
		auth = false
	}
	username := session.Values["username"]
	if username == nil {
		username = ""
	}
	h.user = User{
		IsAuthenticated: auth.(bool),
		Username:        username.(string),
	}
}

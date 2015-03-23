package handlers

import (
	"github.com/gorilla/sessions"

	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

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

type User struct {
	IsAuthenticated bool
	Username        string
}

type Handlers struct {
	sessionStore sessions.Store
	session      *sessions.Session
	user         User
}

func New(store sessions.Store) Handlers {
	h := Handlers{
		sessionStore: store,
	}
	return h
}

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
	}{
		Flashes: flashes,
	}
	t.ExecuteTemplate(w, "base", templateArg)
}

func init() {
	gob.Register(Flash{})
}

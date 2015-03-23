package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func (h *Handlers) DisplayLogin(w http.ResponseWriter, r *http.Request, log *log.Logger) {
	session, err := h.sessionStore.Get(r, "user-management")
	if err != nil {
		log.Print(err)
		http.Error(w, "An error occurred", 500)
	}
	auth := session.Values["authenticated"]
	if auth != nil && auth.(bool) {
		session.AddFlash("Already logged in!")
		session.Save(r, w)
		http.Redirect(w, r, "/", 302)
	}
	t := template.Must(template.ParseFiles(
		"templates/login.tmpl",
		"templates/base.tmpl"))
	t.ExecuteTemplate(w, "base", nil)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request, log *log.Logger) {
	session, err := h.sessionStore.Get(r, "user-management")
	if err != nil {
		log.Print(err)
		http.Error(w, "An error occurred", 500)
	}
	// TODO authenticate user
	session.Values["authenticated"] = true
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := h.sessionStore.Get(r, "user-management")
	if err != nil {
		log.Print(err)
		http.Error(w, "An error occurred", 500)
	}
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func (h *Handlers) DisplayRegister() string {
	return ""
}

func (h *Handlers) Register() string {
	return ""
}

func (h *Handlers) Confirm() string {
	return ""
}

func (h *Handlers) DisplayUser() string {
	return ""
}

func (h *Handlers) FollowUser() string {
	return ""
}

func (h *Handlers) UnfollowUser() string {
	return ""
}

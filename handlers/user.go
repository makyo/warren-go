package handlers

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) DisplayLogin(w http.ResponseWriter, r *http.Request, log *log.Logger) {
	if h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Already logged in!"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	h.render(w, r, "login.tmpl")
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request, log *log.Logger) {
	if h.user.IsAuthenticated {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err.Error())
		http.Error(w, "Could not parse form", http.StatusInternalServerError)
	}
	username, password := r.FormValue("username"), r.FormValue("password")
	hashword, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
	}
	// TODO authenticate against db
	if err := bcrypt.CompareHashAndPassword(hashword, []byte("password")); err != nil {
		h.session.AddFlash(NewFlash("Wrong username or password"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	h.session.Values["authenticated"] = true
	h.session.Values["username"] = username
	h.session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	h.session.Values["authenticated"] = false
	h.session.Values["username"] = nil
	h.session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
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

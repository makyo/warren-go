// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in 
// the LICENSE file.

package handlers

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"

	"github.com/makyo/warren-go/models"
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
	user, err := models.GetUser(username, h.db)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Could not search for user", http.StatusInternalServerError)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Hashword), []byte(password)); err != nil {
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

func (h *Handlers) DisplayRegister(w http.ResponseWriter, r *http.Request) {
	if h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Already logged in!"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	h.render(w, r, "register.tmpl")
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	if h.user.IsAuthenticated {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err.Error())
		http.Error(w, "Could not parse form", http.StatusInternalServerError)
	}
	username, email, password, passwordConfirm := r.FormValue("username"), r.FormValue("email"), r.FormValue("password"), r.FormValue("passwordconfirm")
	if username == "" || email == "" || password == "" {
		h.session.AddFlash(NewFlash("All fields required!", "error"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
	if password != passwordConfirm {
		h.session.AddFlash(NewFlash("Passwords did not match!", "error"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
	c := h.db.C("users")
	existing, err := c.Find(bson.M{"username": username}).Count()
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Could not execute find", http.StatusInternalServerError)
	}
	if existing > 0 {
		h.session.AddFlash(NewFlash("Username taken!", "error"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
	hashword, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Could not generate hashword", http.StatusInternalServerError)
	}
	user := models.User{
		Username: username,
		Email:    email,
		Hashword: string(hashword),
	}
	user.Save(h.db)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *Handlers) Confirm(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handlers) DisplayUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handlers) FollowUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handlers) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handlers) RequestFriendship(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handlers) ConfirmFriendship(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handlers) RejectFriendship(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handlers) CancelFriendship(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

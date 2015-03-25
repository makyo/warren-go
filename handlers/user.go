// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package handlers

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2/bson"

	"github.com/makyo/warren-go/models"
)

func (h *Handlers) DisplayLogin(w http.ResponseWriter, r *http.Request, log *log.Logger, render render.Render) {
	if h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Already logged in!"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	render.HTML(200, "user/login", map[string]interface{}{
		"Title": "Log in",
		"User": h.user,
		"Flashes": h.flashes(r, w),
	})
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request, log *log.Logger) {
	if h.user.IsAuthenticated {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err.Error())
		http.Error(w, "Could not parse form", http.StatusInternalServerError)
		return
	}
	username, password := r.FormValue("username"), r.FormValue("password")
	user, err := models.GetUser(username, h.db)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Could not search for user", http.StatusInternalServerError)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Hashword), []byte(password)); err != nil {
		h.session.AddFlash(NewFlash("Wrong username or password"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
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

func (h *Handlers) DisplayRegister(w http.ResponseWriter, r *http.Request, render render.Render) {
	if h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Already logged in!"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	render.HTML(200, "user/register", map[string]interface{}{
		"Title": "Sign up",
		"User": h.user,
		"Flashes": h.flashes(r, w),
	})
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	if h.user.IsAuthenticated {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err.Error())
		http.Error(w, "Could not parse form", http.StatusInternalServerError)
		return
	}
	username, email, password, passwordConfirm := r.FormValue("username"), r.FormValue("email"), r.FormValue("password"), r.FormValue("passwordconfirm")
	if username == "" || email == "" || password == "" {
		h.session.AddFlash(NewFlash("All fields required!", "warning"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}
	if password != passwordConfirm {
		h.session.AddFlash(NewFlash("Passwords did not match!", "warning"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}
	c := h.db.C("users")
	existing, err := c.Find(bson.M{"username": username}).Count()
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Could not execute find", http.StatusInternalServerError)
		return
	}
	if existing > 0 {
		h.session.AddFlash(NewFlash("Username taken!", "warning"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}
	hashword, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Could not generate hashword", http.StatusInternalServerError)
		return
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

func (h *Handlers) DisplayUser(w http.ResponseWriter, r *http.Request, l *log.Logger, params martini.Params, render render.Render) {
	username := params["username"]
	var user models.User
	if h.user.IsAuthenticated && username == h.user.Model.Username {
		user = h.user.Model
	} else {
		user, err := models.GetUser(username, h.db)
		if err != nil {
			l.Print(err.Error())
			http.Error(w, "Could not fetch user from database", http.StatusInternalServerError)
			return
		}
		if user.Username == "" {
			http.Error(w, "Could not find user", http.StatusNotFound)
			return
		}
	}
	render.HTML(200, "user/displayUser", map[string]interface{}{
		"Title": fmt.Sprintf("User %s", user.Username),
		"User": h.user,
		"Flashes": h.flashes(r, w),
		"DisplayUser": user,
	})
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

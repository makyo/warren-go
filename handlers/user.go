// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2/bson"

	"github.com/warren-community/warren/contenttype/user"
	"github.com/warren-community/warren/models"
)

// Display a login form (or redirect if the user is already logged in).
func (h *Handlers) DisplayLogin(w http.ResponseWriter, r *http.Request, log *log.Logger, render render.Render) {
	if h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Already logged in!"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	render.HTML(http.StatusOK, "user/displayLogin", map[string]interface{}{
		"Title":   "Log in",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
		"CSRF":    h.session.Values["_csrf_token"],
	})
}

// Log the user in.
func (h *Handlers) Login(w http.ResponseWriter, r *http.Request, render render.Render, log *log.Logger) {
	if h.user.IsAuthenticated {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	username, password := r.FormValue("username"), r.FormValue("password")
	user, err := models.GetUser(username, h.db)
	if err != nil {
		log.Printf("Could not get user: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	if !user.Authenticate(password) {
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

// Log the user out.
func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	h.session.Values["authenticated"] = false
	h.session.Values["username"] = nil
	h.session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Display a registration form (or redirect if the user is already logged in).
func (h *Handlers) DisplayRegister(w http.ResponseWriter, r *http.Request, render render.Render) {
	if h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Already logged in!"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	render.HTML(http.StatusOK, "user/displayRegister", map[string]interface{}{
		"Title":   "Sign up",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
		"CSRF":    h.session.Values["_csrf_token"],
	})
}

// Register a new user.
func (h *Handlers) Register(w http.ResponseWriter, r *http.Request, render render.Render) {
	if h.user.IsAuthenticated {
		http.Redirect(w, r, "/", http.StatusSeeOther)
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
		log.Printf("Could not execute find: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	if existing > 0 {
		h.session.AddFlash(NewFlash("Username taken!", "warning"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}
	user, err := models.NewUser(username, email, password)
	if err != nil {
		log.Printf("Could not generate user: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	user.Save(h.db)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// TODO Confirm a user's email address.
func (h *Handlers) Confirm(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Display a user's profile page.
func (h *Handlers) DisplayUser(w http.ResponseWriter, r *http.Request, l *log.Logger, params martini.Params, render render.Render) {
	username := params["username"]
	var user models.User
	if h.user.IsAuthenticated && username == h.user.Model.Username {
		user = h.user.Model
	} else {
		var err error
		user, err = models.GetUser(username, h.db)
		if err != nil {
			l.Printf("Could not fetch user from database: %+v\n", err)
			h.InternalServerError(w, r, render)
			return
		}
		if user.Username == "" {
			h.NotFound(w, r, render)
			return
		}
	}
	entities, err := user.Entities(h.db)
	if err != nil {
		l.Printf("Could not fetch entities from database: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	profile, err := user.Profile(h.db)
	if err != nil {
		if err.Error() == "not found" {
			profile = models.Entity{
				RenderedContent: fmt.Sprintf("%s hasn't created a profile yet.", username),
			}
		} else {
			l.Printf("Could not fetch profile from database: %+v\n", err)
			h.InternalServerError(w, r, render)
			return
		}
	}
	render.HTML(http.StatusOK, "user/displayUser", map[string]interface{}{
		"Title":                  fmt.Sprintf("User %s", user.Username),
		"User":                   h.user,
		"Flashes":                h.flashes(r, w),
		"CSRF":                   h.session.Values["_csrf_token"],
		"DisplayUser":            user,
		"IsFollowing":            h.user.Model.IsFollowing(user.Username),
		"IsFriend":               h.user.Model.IsFriend(user.Username),
		"FriendRequestPending":   h.user.Model.HasRequestedFriendship(user.Username),
		"HasRequestedFriendship": user.HasRequestedFriendship(h.user.Model.Username),
		"Entities":               entities,
		"ProfileStr":             template.HTML(profile.RenderedContent),
	})
}

// Display the page for editing profile and settings.
func (h *Handlers) DisplayEditProfile(w http.ResponseWriter, r *http.Request, l *log.Logger, render render.Render) {
	if !h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Please log in to continue", "warning"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	profile, err := h.user.Model.Profile(h.db)
	if err != nil && err.Error() != "not found" {
		l.Printf("Could not load profile: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	content := user.Profile{}
	if profile.Content != nil {
		content = user.NewProfile(profile.Content.(bson.M))
	}
	render.HTML(http.StatusOK, "user/displayEditProfile", map[string]interface{}{
		"Title":      "Edit profile and settings",
		"User":       h.user,
		"Flashes":    h.flashes(r, w),
		"CSRF":       h.session.Values["_csrf_token"],
		"ProfileStr": content.ProfileText,
		"Pronouns":   content.Pronouns,
		"Website":    content.Website,
	})
}

// Edit the profile information for a user.
func (h *Handlers) EditProfile(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	if !h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Please log in to continue", "warning"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	profileStr, website, pronouns := r.FormValue("profile"), r.FormValue("website"), r.FormValue("pronouns")
	profile, err := h.user.Model.Profile(h.db)
	if err != nil && err.Error() != "not found" {
		l.Printf("Error retrieving existing profile: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	profile.Delete(h.db)
	profile = models.NewEntity(
		"user/profile",
		h.user.Model.Username,
		h.user.Model.Username,
		false,
		"",
		user.Profile{
			ProfileText: profileStr,
			Pronouns:    pronouns,
			Website:     website,
		},
	)
	profile.Save(h.db, h.esConn)
	h.session.AddFlash(NewFlash("Profile updated!", "success"))
	h.session.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/~%s", h.user.Model.Username), http.StatusSeeOther)
}

// Edit the raw settings of a user (password, email).
func (h *Handlers) EditSettings(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	if !h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Please log in to continue", "warning"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	username, password, newpassword, newpasswordconfirm, email := r.FormValue("username"), r.FormValue("password"), r.FormValue("newpassword"), r.FormValue("newpasswordconfirm"), r.FormValue("email")
	if username != h.user.Model.Username {
		l.Print("User attempted to save to another profile")
		h.Forbidden(w, r, render)
		return
	}
	if !h.user.Model.Authenticate(password) {
		h.session.AddFlash(NewFlash("Current password did not match!", "danger"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
		return
	}
	if newpassword != "" {
		if newpassword != newpasswordconfirm {
			h.session.AddFlash(NewFlash("New passwords did not match!", "danger"))
			h.session.Save(r, w)
			http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
			return
		}
		h.user.Model.SetPassword(newpassword)
	}
	h.user.Model.Email = email
	h.user.Model.Save(h.db)
	h.session.AddFlash(NewFlash("Settings updated!", "success"))
	h.session.Save(r, w)
	http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
}

// Attempt to follow a user from the logged-in account.
func (h *Handlers) FollowUser(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	username := r.FormValue("username")
	user, err := models.GetUser(username, h.db)
	if err != nil {
		l.Printf("Could not fetch user from database: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	if user.Username == "" {
		h.NotFound(w, r, render)
		return
	}
	h.user.Model.AddFollowing(&user)
	h.user.Model.Save(h.db)
	user.Save(h.db)
	h.session.AddFlash(NewFlash("User followed!", "success"))
	h.session.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/~%s", username), http.StatusSeeOther)
}

// Attempt to unfollow a user from from the logged-in account.
func (h *Handlers) UnfollowUser(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	username := r.FormValue("username")
	user, err := models.GetUser(username, h.db)
	if err != nil {
		l.Printf("Could not fetch user from database: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	if user.Username == "" {
		h.NotFound(w, r, render)
		return
	}
	h.user.Model.RemoveFollowing(&user)
	h.user.Model.Save(h.db)
	user.Save(h.db)
	h.session.AddFlash(NewFlash("User unfollowed!", "success"))
	h.session.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/~%s", username), http.StatusSeeOther)
}

// Attempt to request a friendship with a user from the logged-in account.
func (h *Handlers) RequestFriendship(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	username := r.FormValue("username")
	user, err := models.GetUser(username, h.db)
	if err != nil {
		l.Printf("Could not fetch user from database: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	if user.Username == "" {
		h.NotFound(w, r, render)
		return
	}
	h.user.Model.RequestFriendship(&user)
	h.user.Model.Save(h.db)
	user.Save(h.db)
	h.session.AddFlash(NewFlash("Friendship requested!", "success"))
	h.session.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/~%s", username), http.StatusSeeOther)
}

// Display currently pending friendship requests for the logged-in account.
func (h *Handlers) DisplayFriendshipRequests(w http.ResponseWriter, r *http.Request, l *log.Logger, render render.Render) {
	if !h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Please log in to continue", "warning"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	render.HTML(http.StatusOK, "user/displayFriendshipRequests", map[string]interface{}{
		"Title":   "Friendship Requests",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
		"CSRF":    h.session.Values["_csrf_token"],
	})
}

// Confirm a friendship request.
func (h *Handlers) ConfirmFriendship(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	username := r.FormValue("username")
	user, err := models.GetUser(username, h.db)
	if err != nil {
		l.Printf("Could not fetch user from database: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	if user.Username == "" {
		h.NotFound(w, r, render)
		return
	}
	h.user.Model.AddFriendship(&user)
	h.user.Model.Save(h.db)
	user.Save(h.db)
	h.session.AddFlash(NewFlash("Friendship confirmed!", "success"))
	h.session.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/~%s", username), http.StatusSeeOther)
}

// Reject a friendship request.
func (h *Handlers) RejectFriendship(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	username := r.FormValue("username")
	user, err := models.GetUser(username, h.db)
	if err != nil {
		l.Printf("Could not fetch user from database: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	if user.Username == "" {
		h.NotFound(w, r, render)
		return
	}
	user.RemoveFriendshipRequest(&h.user.Model)
	h.user.Model.Save(h.db)
	user.Save(h.db)
	h.session.AddFlash(NewFlash("Friendship request rejected!", "success"))
	h.session.Save(r, w)
	http.Redirect(w, r, "/user/friend/requests", http.StatusSeeOther)
}

// Remove a friendship between two accounts.
func (h *Handlers) CancelFriendship(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	username := r.FormValue("username")
	user, err := models.GetUser(username, h.db)
	if err != nil {
		l.Printf("Could not fetch user from database: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	if user.Username == "" {
		h.NotFound(w, r, render)
		return
	}
	h.user.Model.RemoveFriendship(&user)
	h.user.Model.Save(h.db)
	user.Save(h.db)
	h.session.AddFlash(NewFlash("Friendship canceled!", "success"))
	h.session.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/~%s", username), http.StatusSeeOther)
}

package handlers

import (
	"html/template"
	"net/http"

	"github.com/xyproto/permissions2"
)

func DisplayLogin(w http.ResponseWriter) {
	t := template.Must(template.ParseFiles(
		"templates/login.tmpl",
		"templates/base.tmpl"))
	t.ExecuteTemplate(w, "base", nil)
}

func Login(w http.ResponseWriter, req *http.Request, userstate *permissions.UserState) {
	if userstate.IsLoggedIn(userstate.Username(req)) {
		http.Redirect(w, req, "/", 302)
	} else {
		err := req.ParseForm()
		if err != nil {
			http.Error(w, "Problem parsing form", 500)
		}
		username := req.FormValue("username")
		password := req.FormValue("password")
		if err = permissions.ValidUsernamePassword(username, password); err != nil {
			http.Redirect(w, req, "/login", 302)
		} else {
			userstate.Login(w, username)
			http.Redirect(w, req, "/", 302)
		}
		http.Redirect(w, req, "/login", 302)
	}
}

func Logout(w http.ResponseWriter, req *http.Request, userstate *permissions.UserState) string {
	userstate.Logout(userstate.Username(req))
	http.Redirect(w, req, "/", 302)
}

func DisplayRegister() string {
	return ""
}

func Register() string {
	return ""
}

func Confirm() string {
	return ""
}

func DisplayUser() string {
	return ""
}

func FollowUser() string {
	return ""
}

func UnfollowUser() string {
	return ""
}

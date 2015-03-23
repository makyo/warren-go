package handlers

import (
	"html/template"
	"net/http"
)

func DisplayLogin(w http.ResponseWriter) {
	t := template.Must(template.ParseFiles(
		"templates/login.tmpl",
		"templates/base.tmpl"))
	t.ExecuteTemplate(w, "base", nil)
}

func Login() string {
	return ""
}

func Logout() string {
	return ""
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

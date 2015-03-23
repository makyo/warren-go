package main

import (
	"github.com/go-martini/martini"
	"github.com/gorilla/sessions"

	"github.com/makyo/warren-go/handlers"
)

var store = sessions.NewCookieStore([]byte("dev-secret"))

func main() {
	m := martini.Classic()

	h := handlers.New(store)

	m.Get("/", h.Front)

	m.Get("/login", h.DisplayLogin)
	m.Post("/login", h.Login)
	m.Get("/logout", h.Logout)
	m.Get("/register", h.DisplayRegister)
	m.Post("/register", h.Register)
	m.Get("/confirm/:confirmation", h.Confirm)
	m.Get("/~:username", h.DisplayUser)
	m.Post("/~:username/follow", h.FollowUser)
	m.Post("/~:username/unfollow", h.UnfollowUser)

	m.Get("/(?P<post>\\d+)", h.DisplayPost)
	m.Get("/(?P<post>\\d+)/delete", h.DisplayDeletePost)
	m.Get("/(?P<post>\\d+)/delete", h.DeletePost)
	m.Post("/(?P<post>\\d+)/share", h.SharePost)
	m.Get("/post", h.DisplayCreatePost)
	m.Post("/post", h.CreatePost)

	m.Run()
}

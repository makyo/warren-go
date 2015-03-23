package main

import (
	"github.com/go-martini/martini"

	"github.com/makyo/warren-go/handlers"
)

func main() {
	m := martini.Classic()

	m.Get("/", handlers.Front)

	m.Get("/login", handlers.DisplayLogin)
	m.Post("/login", handlers.Login)
	m.Get("/logout", handlers.Logout)
	m.Get("/register", handlers.DisplayRegister)
	m.Post("/register", handlers.Register)
	m.Get("/confirm/:confirmation", handlers.Confirm)
	m.Get("/~:username", handlers.DisplayUser)
	m.Post("/~:username/follow", handlers.FollowUser)
	m.Post("/~:username/unfollow", handlers.UnfollowUser)

	m.Get("/(?P<post>\\d+)", handlers.DisplayPost)
	m.Get("/(?P<post>\\d+)/delete", handlers.DisplayDeletePost)
	m.Get("/(?P<post>\\d+)/delete", handlers.DeletePost)
	m.Post("/(?P<post>\\d+)/share", handlers.SharePost)
	m.Get("/post", handlers.DisplayCreatePost)
	m.Post("/post", handlers.CreatePost)

	m.Run()
}

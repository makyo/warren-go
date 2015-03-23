package handlers

import (
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
)

type User struct {
	IsAuthenticated bool
	Username        string
}

type Handlers struct {
	sessionStore sessions.Store
	session      *sessions.Session
	db           *mgo.Database
	user         User
}

func New(store sessions.Store, db *mgo.Database) Handlers {
	h := Handlers{
		sessionStore: store,
		db:           db,
	}
	return h
}

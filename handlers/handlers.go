package handlers

import (
	"github.com/gorilla/sessions"
)

type Handlers struct {
	sessionStore sessions.Store
}

func New(store sessions.Store) Handlers {
	h := Handlers{
		sessionStore: store,
	}
	return h
}

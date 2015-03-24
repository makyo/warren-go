// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in 
// the LICENSE file.

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

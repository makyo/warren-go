// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package handlers

import (
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"

	"github.com/makyo/warren-go/models"
	elastigo "github.com/mattbaird/elastigo/lib"
)

// A user object to store in the session, containing a User model and a flag
// indicating whether or not the user is logged in.
type User struct {
	IsAuthenticated bool
	Model           models.User
}

// A struct maintaining various bits of information that might be needed within
// the handlers that may not necessarily need to be passed in through the
// service infrastructure.
type Handlers struct {
	sessionStore sessions.Store
	session      *sessions.Session
	db           *mgo.Database
	esConn       *elastigo.Conn
	user         User
}

// Create a new handler object with provided attributes.
func New(store sessions.Store, db *mgo.Database, esConn *elastigo.Conn) Handlers {
	h := Handlers{
		sessionStore: store,
		db:           db,
		esConn:       esConn,
	}
	return h
}

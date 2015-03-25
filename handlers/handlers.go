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

type User struct {
	IsAuthenticated bool
	Model           models.User
}

type Handlers struct {
	sessionStore sessions.Store
	session      *sessions.Session
	db           *mgo.Database
	esConn       *elastigo.Conn
	user         User
}

func New(store sessions.Store, db *mgo.Database, esConn *elastigo.Conn) Handlers {
	h := Handlers{
		sessionStore: store,
		db:           db,
		esConn:       esConn,
	}
	return h
}

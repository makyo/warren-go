// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in 
// the LICENSE file.

package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username  string
	Email     string
	Hashword  string
	Following []string
	Followers []string
	Friends   []string
}

func (u *User) Save(db *mgo.Database) error {
	c := db.C("users")
	_, err := c.Upsert(bson.M{"username": u.Username}, u)
	return err
}

func GetUser(username string, db *mgo.Database) (User, error) {
	user := User{}
	c := db.C("users")
	err := c.Find(bson.M{"username": username}).One(&user)
	return user, err
}

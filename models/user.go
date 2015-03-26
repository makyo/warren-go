// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package models

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// A representation of a user in the database, storing only minimal
// information.
type User struct {
	Username             string
	Email                string
	Hashword             []byte
	Following            []string
	Followers            []string
	Friends              []string
	FriendRequests       []string
	FriendshipsRequested []string
}

// Retrieve a user from the database given a username
func GetUser(username string, db *mgo.Database) (User, error) {
	var user User
	q := db.C("users").Find(bson.M{"username": username})
	if c, err := q.Count(); c == 0 {
		return user, err
	}
	err := q.One(&user)
	return user, err
}

func NewUser(username string, email string, password string) (User, error) {
	hashword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return User{}, err
	}
	return User{
		Username: username,
		Email:    email,
		Hashword: hashword,
	}, nil
}

// Save a given user model to the database.
func (u *User) Save(db *mgo.Database) error {
	_, err := db.C("users").Upsert(bson.M{"username": u.Username}, u)
	return err
}

// Attempt to authenticate user with a password.
func (u *User) Authenticate(password string) bool {
	if err := bcrypt.CompareHashAndPassword(u.Hashword, []byte(password)); err != nil {
		return false
	}
	return true
}

// Fetch all given entities for a user
// XXX This is a naive and expensive approach
func (u *User) Entities(db *mgo.Database) ([]Entity, error) {
	var result []Entity
	q := db.C("entities").Find(bson.M{"owner": u.Username})
	err := q.All(&result)
	return result, err
}

// Return true if the user is following another user by that username.
func (u *User) IsFollowing(username string) bool {
	for _, name := range u.Following {
		if name == username {
			return true
		}
	}
	return false
}

// Add a unidirectional following relationship between two users
func (u *User) AddFollowing(user *User) {
	if !u.IsFollowing(user.Username) {
		u.Following = append(u.Following, user.Username)
		user.Followers = append(user.Followers, u.Username)
	}
}

// Remove a following relationship between two users.
func (u *User) RemoveFollowing(user *User) {
	if u.IsFollowing(user.Username) {
		for i, name := range u.Following {
			if name == user.Username {
				u.Following = append(u.Following[:i], u.Following[i+1:]...)
				break
			}
		}
		for i, name := range user.Followers {
			if name == u.Username {
				user.Followers = append(user.Followers[:i], user.Followers[i+1:]...)
				break
			}
		}
	}
}

// Return true of the user is friends with the user by the given username.
func (u *User) IsFriend(username string) bool {
	for _, name := range u.Friends {
		if name == username {
			return true
		}
	}
	return false
}

// Return true if the user has requested to be friends with the user by
// the given username.
func (u *User) HasRequestedFriendship(username string) bool {
	for _, name := range u.FriendshipsRequested {
		if name == username {
			return true
		}
	}
	return false
}

// Create a friendship request between this user and the given user.
func (u *User) RequestFriendship(user *User) {
	if !u.HasRequestedFriendship(user.Username) && !u.IsFriend(user.Username) {
		u.FriendshipsRequested = append(u.FriendshipsRequested, user.Username)
		user.FriendRequests = append(user.FriendRequests, u.Username)
	}
}

// Remove a friendship request between this user and the given user.
func (u *User) RemoveFriendshipRequest(user *User) {
	if u.HasRequestedFriendship(user.Username) {
		for i, name := range user.FriendRequests {
			if name == u.Username {
				user.FriendRequests = append(user.FriendRequests[:i], user.FriendRequests[i+1:]...)
				break
			}
		}
		for i, name := range u.FriendshipsRequested {
			if name == user.Username {
				u.FriendshipsRequested = append(u.FriendshipsRequested[:i], u.FriendshipsRequested[i+1:]...)
				break
			}
		}
	}
}

// Add a bidirectional friendship relationship between the two users, removing
// any pending requests if they exist.
func (u *User) AddFriendship(user *User) {
	if !u.IsFriend(user.Username) {
		u.Friends = append(u.Friends, user.Username)
		user.Friends = append(user.Friends, u.Username)
		if u.HasRequestedFriendship(user.Username) {
			u.RemoveFriendshipRequest(user)
		}
		if user.HasRequestedFriendship(u.Username) {
			user.RemoveFriendshipRequest(u)
		}
	}
}

// Remove a friendship relationship between the two users.
func (u *User) RemoveFriendship(user *User) {
	if u.IsFriend(user.Username) {
		for i, name := range u.Friends {
			if name == user.Username {
				u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
				break
			}
		}
		for i, name := range user.Friends {
			if name == u.Username {
				user.Friends = append(user.Friends[:i], user.Friends[i+1:]...)
				break
			}
		}
	}
}

// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username             string
	Email                string
	Hashword             string
	Following            []string
	Followers            []string
	Friends              []string
	FriendRequests       []string
	FriendshipsRequested []string
}

func (u *User) Save(db *mgo.Database) error {
	c := db.C("users")
	_, err := c.Upsert(bson.M{"username": u.Username}, u)
	return err
}

func (u *User) IsFollowing(username string) bool {
	for _, name := range u.Following {
		if name == username {
			return true
		}
	}
	return false
}

func (u *User) AddFollowing(user *User) {
	if !u.IsFollowing(user.Username) {
		u.Following = append(u.Following, user.Username)
		user.Followers = append(user.Followers, u.Username)
	}
}

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

func (u *User) IsFriend(username string) bool {
	for _, name := range u.Friends {
		if name == username {
			return true
		}
	}
	return false
}

func (u *User) HasRequestedFriendship(username string) bool {
	for _, name := range u.FriendshipsRequested {
		if name == username {
			return true
		}
	}
	return false
}

func (u *User) RequestFriendship(user *User) {
	if !u.HasRequestedFriendship(user.Username) {
		u.FriendshipsRequested = append(u.FriendshipsRequested, user.Username)
		user.FriendRequests = append(user.FriendRequests, u.Username)
	}
}

func (u *User) RemoveFriendshipRequest(user *User) {
	if user.HasRequestedFriendship(u.Username) {
		fmt.Println("Found a request")
		for i, name := range u.FriendRequests {
			if name == user.Username {
				u.FriendRequests = append(u.FriendRequests[:i], u.FriendRequests[i+1:]...)
				break
			}
		}
		for i, name := range user.FriendshipsRequested {
			if name == u.Username {
				user.FriendshipsRequested = append(user.FriendshipsRequested[:i], user.FriendshipsRequested[i+1:]...)
				break
			}
		}
	}
}

func (u *User) AddFriendship(user *User) {
	if !u.IsFriend(user.Username) {
		u.Friends = append(u.Friends, user.Username)
		user.Friends = append(user.Friends, u.Username)
		u.RemoveFriendshipRequest(user)
		if user.HasRequestedFriendship(u.Username) {
			user.RemoveFriendshipRequest(u)
		}
	}
}

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

func GetUser(username string, db *mgo.Database) (User, error) {
	var user User
	c := db.C("users")
	q := c.Find(bson.M{"username": username})
	if c, err := q.Count(); c == 0 {
		return user, err
	}
	err := q.One(&user)
	return user, err
}

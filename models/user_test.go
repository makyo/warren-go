package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserModel(t *testing.T) {
	Convey("Given two users", t, func() {
		a, b := User{Username: "a"}, User{Username: "b"}

		Convey("When one follows the other", func() {
			a.AddFollowing(&b)

			Convey("IsFollowing should show the unidirectional relation", func() {
				Convey("It should be unidirectional between A and B", func() {
					So(a.IsFollowing("b"), ShouldBeTrue)
					So(b.IsFollowing("a"), ShouldBeFalse)
				})
			})

			Convey("AddFollowing should be unidirectional", func() {
				Convey("Directionality should be evident", func() {
					So(a.Following, ShouldResemble, []string{"b"})
					So(b.Followers, ShouldResemble, []string{"a"})
					So(a.Followers, ShouldBeNil)
					So(b.Following, ShouldBeNil)
				})
			})

			Convey("When removing followers", func() {
				Convey("Should not panic removing non-following users", func() {
					c := User{Username: "c"}
					a.RemoveFollowing(&c)
				})

				Convey("Will remove following for both users", func() {
					a.RemoveFollowing(&b)

					Convey("Removing should work on both A and B", func() {
						So(a.Followers, ShouldBeEmpty)
						So(a.Following, ShouldBeEmpty)
						So(b.Followers, ShouldBeEmpty)
						So(b.Following, ShouldBeEmpty)
					})
				})
			})
		})

		Convey("Who are friends, and two other users who want to be", func() {
			c, d := User{Username: "c"}, User{Username: "d"}
			a.AddFriendship(&b)
			c.RequestFriendship(&d)

			Convey("IsFriend should show the bidirectional relation", func() {
				Convey("True between A and B", func() {
					So(a.IsFriend("b"), ShouldBeTrue)
					So(b.IsFriend("a"), ShouldBeTrue)
				})

				Convey("False between non-friends", func() {
					So(a.IsFriend("c"), ShouldBeFalse)
				})

				Convey("False between pending friends", func() {
					So(c.IsFriend("d"), ShouldBeFalse)
				})
			})

			Convey("HasRequestedFriendship should show the bidirectional relation", func() {
				Convey("It should be directional", func() {
					So(c.HasRequestedFriendship("d"), ShouldBeTrue)
					So(d.HasRequestedFriendship("c"), ShouldBeFalse)
				})

				Convey("It should be alse for friends", func() {
					So(a.HasRequestedFriendship("b"), ShouldBeFalse)
				})

				Convey("It should be false for non-friends", func() {
					So(a.HasRequestedFriendship("c"), ShouldBeFalse)
				})
			})

			Convey("RequestFriendship should create a requested relation", func() {
				e := User{Username: "e"}

				Convey("Adds a friendship request", func() {
					e.RequestFriendship(&a)

					So(e.HasRequestedFriendship("a"), ShouldBeTrue)
				})

				Convey("Ignores friends", func() {
					a.RequestFriendship(&b)
					So(a.HasRequestedFriendship("b"), ShouldBeFalse)
				})

				Convey("Does not duplicate requests", func() {
					So(c.FriendshipsRequested, ShouldResemble, []string{"d"})
					So(d.FriendRequests, ShouldResemble, []string{"c"})
					c.RequestFriendship(&d)
					So(c.FriendshipsRequested, ShouldResemble, []string{"d"})
					So(d.FriendRequests, ShouldResemble, []string{"c"})
				})
				
				Convey("Allows requests going the other way", func() {
					d.RequestFriendship(&c)
					So(d.HasRequestedFriendship("c"), ShouldBeTrue)
				})
			})

			Convey("RemoveFriendshipRequest should remove a request", func() {
				So(a.HasRequestedFriendship("b"), ShouldBeFalse)
				So(c.HasRequestedFriendship("d"), ShouldBeTrue)
				So(d.HasRequestedFriendship("c"), ShouldBeFalse)

				Convey("Ignores users without a request", func() {
					a.RemoveFriendshipRequest(&b)
					So(a.HasRequestedFriendship("b"), ShouldBeFalse)
				})

				Convey("Ignores inverse", func() {
					d.RemoveFriendshipRequest(&c)
					So(d.HasRequestedFriendship("c"), ShouldBeFalse)
					So(c.HasRequestedFriendship("d"), ShouldBeTrue)
				})

				Convey("Removes friendship request", func() {
					c.RemoveFriendshipRequest(&d)
					So(c.HasRequestedFriendship("d"), ShouldBeFalse)
					So(c.FriendshipsRequested, ShouldResemble, []string{})
					So(d.FriendRequests, ShouldResemble, []string{})
				})
			})

			Convey("AddFriendship should confirm a friendship", func() {
				Convey("Adds a friendship", func() {
					c.AddFriendship(&d)
					So(c.IsFriend("d"), ShouldBeTrue)
					So(d.IsFriend("c"), ShouldBeTrue)
					So(c.HasRequestedFriendship("d"), ShouldBeFalse)
				})

				Convey("Ignores friends", func() {
					a.AddFriendship(&b)
					So(a.IsFriend("b"), ShouldBeTrue)
					So(b.IsFriend("a"), ShouldBeTrue)
					So(a.Friends, ShouldResemble, []string{"b"})
					So(b.Friends, ShouldResemble, []string{"a"})
				})
			})

			Convey("RemoveFriendship should remove a friendship", func() {
				Convey("Removes a friendship", func() {
					a.RemoveFriendship(&b)
					So(a.IsFriend("b"), ShouldBeFalse)
					So(b.IsFriend("a"), ShouldBeFalse)
				})

				Convey("Ignores non-friends", func() {
					c.RemoveFriendship(&d)
					So(c.IsFriend("d"), ShouldBeFalse)
					So(d.IsFriend("c"), ShouldBeFalse)
				})
			})
		})
	})
}
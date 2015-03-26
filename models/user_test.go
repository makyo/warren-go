package models

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type UserModelFollowingSuite struct {
	a User
	b User
}

func (s *UserModelFollowingSuite) SetUpTest(c *C) {
	s.a, s.b = User{Username: "a"}, User{Username: "b"}
	s.a.AddFollowing(&s.b)
}

var _ = Suite(&UserModelFollowingSuite{})

func (s *UserModelFollowingSuite) TestIsFollowing(c *C) {
	// IsFollowing returns true only if the user is following username.
	c.Assert(s.a.IsFollowing("b"), Equals, true)
	c.Assert(s.a.IsFollowing("c"), Equals, false)
}

func (s *UserModelFollowingSuite) TestAddFollowing(c *C) {
	// Following shows for both users.
	c.Assert(s.a.Following, DeepEquals, []string{"b"})
	c.Assert(s.b.Followers, DeepEquals, []string{"a"})

	// Following is unidirectional.
	c.Assert(s.a.Followers, IsNil)
	c.Assert(s.b.Following, IsNil)
}

func (s *UserModelFollowingSuite) TestRemoveFollowing(c *C) {
	// Will not panic with non-following users
	user := User{Username: "c"}
	s.a.RemoveFollowing(&user)

	// Will remove following for both users.
	s.a.RemoveFollowing(&s.b)
	c.Assert(s.a.Followers, HasLen, 0)
	c.Assert(s.a.Following, HasLen, 0)
	c.Assert(s.b.Followers, HasLen, 0)
	c.Assert(s.b.Following, HasLen, 0)
}

type UserModelFriendSuite struct {
	a User
	b User
	c User
	d User
}

func (s *UserModelFriendSuite) SetUpTest(c *C) {
	s.a, s.b, s.c, s.d = User{Username: "a"}, User{Username: "b"}, User{Username: "c"}, User{Username: "d"}
	s.a.AddFriendship(&s.b)
	s.c.RequestFriendship(&s.d)
}

var _ = Suite(&UserModelFollowingSuite{})

func (s *UserModelFriendSuite) TestIsFriend(c *C) {
	// True for friends
	c.Assert(s.a.IsFriend("b"), Equals, true)

	// False for non-friends
	c.Assert(s.a.IsFriend("c"), Equals, false)

	// False for pending friendship
	c.Assert(s.c.IsFriend("d"), Equals, false)
}

func (s *UserModelFriendSuite) TestHasRequestedFriendship(c *C) {
	// True when one has requested friendship of another
	c.Assert(s.c.HasRequestedFriendship("d"), Equals, true)

	// False in the other direction
	c.Assert(s.d.HasRequestedFriendship("c"), Equals, false)

	// False for friends
	c.Assert(s.a.HasRequestedFriendship("b"), Equals, false)

	// False for non-friends
	c.Assert(s.a.HasRequestedFriendship("c"), Equals, false)
}

func (s *UserModelFriendSuite) TestRequestFriendship(c *C) {
	user := User{Username: "e"}

	// Adds a friendship request
	user.RequestFriendship(&s.a)
	c.Assert(user.HasRequestedFriendship("a"), Equals, true)

	// Ignores friends
	s.a.RequestFriendship(&s.b)
	c.Assert(user.HasRequestedFriendship("b"), Equals, false)

	// Does not duplicate requests
	c.Assert(s.c.FriendshipsRequested, DeepEquals, []string{"d"})
	c.Assert(s.d.FriendRequests, DeepEquals, []string{"c"})
	s.c.RequestFriendship(&s.d)
	c.Assert(s.c.FriendshipsRequested, DeepEquals, []string{"d"})
	c.Assert(s.d.FriendRequests, DeepEquals, []string{"c"})

	// Allows requests going the other way
	s.a.RequestFriendship(&user)
	c.Assert(s.a.HasRequestedFriendship("e"), Equals, true)
}

func (s *UserModelFriendSuite) TestRemoveFriendshipRequest(c *C) {
	// Ignores non-friends
	s.a.RemoveFriendshipRequest(&s.b)
	c.Assert(s.a.HasRequestedFriendship("b"), Equals, false)

	// Ignores inverse
	s.d.RemoveFriendshipRequest(&s.c)
	c.Assert(s.d.HasRequestedFriendship("c"), Equals, false)
	c.Assert(s.c.HasRequestedFriendship("d"), Equals, true)

	// Removes friendship request
	s.c.RemoveFriendshipRequest(&s.d)
	c.Assert(s.c.HasRequestedFriendship("d"), Equals, false)
	c.Assert(s.c.FriendshipsRequested, DeepEquals, []string{})
	c.Assert(s.d.FriendRequests, DeepEquals, []string{})
}

func (s *UserModelFriendSuite) TestAddFriendship(c *C) {
	// Adds friendship
	s.c.AddFriendship(&s.d)
	c.Assert(s.c.IsFriend("d"), Equals, true)
	c.Assert(s.d.IsFriend("c"), Equals, true)
	c.Assert(s.c.HasRequestedFriendship("d"), Equals, false)

	// Ignores friends
	s.a.AddFriendship(&s.b)
	c.Assert(s.a.IsFriend("b"), Equals, true)
	c.Assert(s.b.IsFriend("a"), Equals, true)
}

func (s *UserModelFriendSuite) TestRemoveFriendship(c *C) {
	// Removes friendship
	s.a.RemoveFriendship(&s.b)
	c.Assert(s.a.IsFriend("b"), Equals, false)
	c.Assert(s.b.IsFriend("a"), Equals, false)

	// Ignores non-friends
	s.c.RemoveFriendship(&s.d)
	c.Assert(s.c.IsFriend("d"), Equals, false)
	c.Assert(s.d.IsFriend("c"), Equals, false)
}

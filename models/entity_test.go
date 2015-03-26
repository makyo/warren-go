// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEntityModel(t *testing.T) {
	Convey("Given an entity", t, func() {
		e := NewEntity("type", "owner", "originalOwner", true, "title", "content")

		Convey("The fields are created properly", func() {
			So(e.Id.Valid(), ShouldBeTrue)
			So(e.ContentType, ShouldEqual, "type")
			So(e.Owner, ShouldEqual, "owner")
			So(e.OriginalOwner, ShouldEqual, "originalOwner")
			So(e.IsShare, ShouldBeTrue)
			So(e.Title, ShouldEqual, "title")
			So(e.Content, ShouldEqual, "content")
		})

		Convey("The rendered content can be created", func() {
			e.updateRenderedContent()
			So(e.RenderedContent, ShouldEqual, "content")
		})

		Convey("The indexed content can be created", func() {
			e.updateIndexedContent()
			So(e.IndexedContent, ShouldEqual, "content")
		})

		Convey("Ownership can be asserted", func() {
			owner := User{Username: "owner"}
			notOwner := User{Username: "bad-wolf"}
			So(e.BelongsToUser(owner), ShouldBeTrue)
			So(e.BelongsToUser(notOwner), ShouldBeFalse)
		})
	})
}

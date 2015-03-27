// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package models

import (
	"testing"

	"github.com/warren-community/warren/contenttype"
	"github.com/warren-community/warren/contenttype/text"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEntityModel(t *testing.T) {
	Convey("Given an entity", t, func() {
		contenttype.Registry["text/raw"] = new(text.Raw)
		e := NewEntity("text/raw", "owner", "originalOwner", true, "title", "content")

		Convey("The fields are created properly", func() {
			So(e.Id.Valid(), ShouldBeTrue)
			So(e.ContentType, ShouldEqual, "text/raw")
			So(e.Owner, ShouldEqual, "owner")
			So(e.OriginalOwner, ShouldEqual, "originalOwner")
			So(e.IsShare, ShouldBeTrue)
			So(e.Title, ShouldEqual, "title")
			So(e.Content, ShouldEqual, "content")
		})

		Convey("The rendered content can be created", func() {
			err := e.updateRenderedContent(true)
			So(e.RenderedContent, ShouldEqual, "content")
			So(err, ShouldBeNil)
		})

		Convey("Rendering content checks for safety", func() {
			err := e.updateRenderedContent(false)
			So(e.RenderedContent, ShouldEqual, "")
			So(err.Error(), ShouldResemble, "Attempted unsafe content-type usage: text/raw")
			err = e.updateIndexedContent(false)
			So(e.IndexedContent, ShouldEqual, "")
			So(err.Error(), ShouldResemble, "Attempted unsafe content-type usage: text/raw")
		})

		Convey("The indexed content can be created", func() {
			err := e.updateIndexedContent(true)
			So(e.IndexedContent, ShouldEqual, "content")
			So(err, ShouldBeNil)
		})

		Convey("Ownership can be asserted", func() {
			owner := User{Username: "owner"}
			notOwner := User{Username: "bad-wolf"}
			So(e.BelongsToUser(owner), ShouldBeTrue)
			So(e.BelongsToUser(notOwner), ShouldBeFalse)
		})
	})
}

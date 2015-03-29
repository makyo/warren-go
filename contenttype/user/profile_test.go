// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package user

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserProfileContentType(t *testing.T) {
	Convey("The user profile content type", t, func() {
		ct := Profile{
			ProfileText: "# profile text",
			Pronouns:    "pronouns",
			Website:     "website",
		}

		Convey("Should be able to be created from a bson map", func() {
			ct = NewProfile(bson.M{
				"profiletext": "foo",
				"pronouns":    "bar",
				"website":     "baz",
			})
			So(ct, ShouldResemble, Profile{
				ProfileText: "foo",
				Pronouns:    "bar",
				Website:     "baz",
			})

			Convey("Blank fields can be allowed", func() {
				ct = NewProfile(bson.M{})
				So(ct, ShouldResemble, Profile{})
			})

			Convey("Unknown fields are ignored", func() {
				ct = NewProfile(bson.M{"asdf": "bad-wolf"})
				So(ct, ShouldResemble, Profile{})
			})
		})

		Convey("Should be safe", func() {
			So(ct.Safe(), ShouldBeTrue)
		})

		Convey("Renders profile for display", func() {
			result, err := ct.RenderDisplayContent(ct)
			So(err, ShouldBeNil)
			So(result, ShouldEqual, `<div class="well well-sm">
	<h1>profile text</h1>

</div>
<dl>
	
		<dt>Pronouns</dt>
		<dd>pronouns</dd>
	
	<dt>Website</dt>
	<dd>
		
			<a href="website" target="_blank">website <span class="glyphicon glyphicon-share" aria-hidden="true"></span></a>
		
	</dd>
</dl>`)
		})

		Convey("Renders profile for indexing", func() {
			result, err := ct.RenderIndexContent(ct)
			So(err, ShouldBeNil)
			So(result, ShouldEqual, ct.ProfileText)
		})
	})
}

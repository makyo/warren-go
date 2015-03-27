// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package user

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserProfileContentType(t *testing.T) {
	Convey("The user profile content type", t, func() {
		ct := new(Profile)

		Convey("Should be safe", func() {
			So(ct.Safe(), ShouldBeTrue)
		})

		SkipConvey("Renders profile for display", func() {
			// Not Implemented
		})

		SkipConvey("Renders profile for indexing", func() {
			// Not Implemented
		})
	})
}

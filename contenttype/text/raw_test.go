// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package text

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTextRawContentType(t *testing.T) {
	Convey("The text/raw content type", t, func() {
		ct := new(Raw)

		Convey("Should be unsafe", func() {
			So(ct.Safe(), ShouldBeFalse)
		})

		Convey("Should not modify content for index or rendering", func() {
			content, err := ct.RenderDisplayContent("content<script></script>")
			So(content, ShouldEqual, "content<script></script>")
			So(err, ShouldBeNil)

			content, err = ct.RenderIndexContent("content<script></script>")
			So(content, ShouldEqual, "content<script></script>")
			So(err, ShouldBeNil)
		})
	})
}

// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package text

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTextPlainContentType(t *testing.T) {
	Convey("The text/plain content type", t, func() {
		ct := new(Plain)

		Convey("Should be safe", func() {
			So(ct.Safe(), ShouldBeTrue)
		})

		Convey("Should sanitize html for display", func() {
			content, err := ct.RenderDisplayContent(`<img src="foo" />`)
			So(content, ShouldEqual, `<pre>&lt;img src=&#34;foo&#34; /&gt;</pre>`)
			So(err, ShouldBeNil)
		})

		Convey("Should add linebreaks for display", func() {
			content, err := ct.RenderDisplayContent(`bad

wolf
42`)
			So(content, ShouldEqual, `<pre>bad

wolf
42</pre>`)
			So(err, ShouldBeNil)
		})

		Convey("Should not modify for indexing", func() {
			content, err := ct.RenderIndexContent("<img />foo")
			So(content, ShouldEqual, "<img />foo")
			So(err, ShouldBeNil)
		})
	})
}

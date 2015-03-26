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
			So(content, ShouldEqual, `<p>&lt;img src=&#34;foo&#34; /&gt;</p>`)
			So(err, ShouldBeNil)
		})

		Convey("Should add linebreaks for display", func() {
			content, err := ct.RenderDisplayContent("bad\r\n\r\nwolf\r\n42")
			So(content, ShouldEqual, `<p>bad</p><p>wolf<br />42</p>`)
			So(err, ShouldBeNil)
		})

		Convey("Should not modify for indexing", func() {
			content, err := ct.RenderIndexContent("<img />foo")
			So(content, ShouldEqual, "<img />foo")
			So(err, ShouldBeNil)
		})
	})
}

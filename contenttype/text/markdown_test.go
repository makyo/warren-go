// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package text

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTextMarkdownContentType(t *testing.T) {
	Convey("The text/markdown content type", t, func() {
		ct := new(Markdown)

		Convey("Should be safe", func() {
			So(ct.Safe(), ShouldBeTrue)
		})

		Convey("Should render markdown for display", func() {
			// We don't need to test blackfriday, just make sure that it's
			// being called on the content.
			given := `
# foo

* a
* list`
			expected := `<h1>foo</h1>

<ul>
<li>a</li>
<li>list</li>
</ul>
`
			content, err := ct.RenderDisplayContent(given)
			So(content, ShouldEqual, expected)
			So(err, ShouldBeNil)
		})

		Convey("Should sanitize html for display", func() {
			// It's blackfriday's fault that the <p> tag isn't closed, not ours.
			// We don't need to test bluemonday, just make sure that it's
			// being called on the content.
			content, err := ct.RenderDisplayContent(`<a href="javascript:alert('XSS1')" onmouseover="alert('XSS2')">XSS<a>`)
			So(content, ShouldEqual, "<p>XSS\n")
			So(err, ShouldBeNil)
		})

		Convey("Should leave markdown alone for indexing", func() {
			content, err := ct.RenderIndexContent("# foo")
			So(content, ShouldEqual, "# foo")
			So(err, ShouldBeNil)
		})
	})
}

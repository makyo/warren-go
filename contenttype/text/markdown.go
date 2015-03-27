// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package text

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

// The text/markdown content type.
type Markdown struct{}

// Since the content is sanitized, this content type is safe.
func (c *Markdown) Safe() bool {
	return true
}

// Render the markdown, sanitize the output, and return that for display.
func (c *Markdown) RenderDisplayContent(content string) (string, error) {
	rendered := blackfriday.MarkdownCommon([]byte(content))
	safe := bluemonday.UGCPolicy().SanitizeBytes(rendered)
	return string(safe), nil
}

// Simply return the markdown for indexing.
func (c *Markdown) RenderIndexContent(content string) (string, error) {
	return content, nil
}

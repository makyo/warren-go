// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package text

import (
	"fmt"
	"html/template"
	"regexp"
)

// The text/plain content type.
type Plain struct{}

// Since the display content is sanitized, this content type is safe.
func (c *Plain) Safe() bool {
	return true
}

// Sanitize the output, replace newlines with HTML line breaks, and return
// the modified content.
func (c *Plain) RenderDisplayContent(content interface{}) (string, error) {
	contentStr := template.HTMLEscapeString(content.(string))
	paraRe, err := regexp.Compile("\r?\n\r?\n")
	if err != nil {
		return "", err
	}
	breakRe, err := regexp.Compile("\r?\n")
	if err != nil {
		return "", err
	}
	content = string(paraRe.ReplaceAll([]byte(contentStr), []byte("</p><p>")))
	content = string(breakRe.ReplaceAll([]byte(contentStr), []byte("<br />")))
	return fmt.Sprintf("<p>%s</p>", contentStr), nil
}

// Simply return the content.
func (c *Plain) RenderIndexContent(content interface{}) (string, error) {
	return content.(string), nil
}

// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package text

// The text/raw content type.
type Raw struct{}

// This should only ever be used for testing, as it does not sanitize output in
// any way.
func (c *Raw) Safe() bool {
	return false
}

// Simply returns the content.
func (c *Raw) RenderDisplayContent(content interface{}) (string, error) {
	return content.(string), nil
}

// Simply returns the content.
func (c *Raw) RenderIndexContent(content interface{}) (string, error) {
	return content.(string), nil
}

// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package contenttype

import (
	"github.com/warren-community/warren/contenttype/test"
	"github.com/warren-community/warren/contenttype/text"
	"github.com/warren-community/warren/contenttype/user"
)

// Default content type to be used when no content type is specified or no
// matching content type is found in the registry.
var DefaultContentType = new(text.Plain)

// The registry maps content type strings to instances of ContentType.
// XXX This is probably less than ideal, and could maybe be done with
// introspection, but will serve for now.
var Registry = map[string]ContentType{
	"text/plain":    new(text.Plain),
	"text/markdown": new(text.Markdown),
	"user/profile":  new(user.Profile),
	"test/error":    new(test.Error),
}

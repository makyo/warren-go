package contenttype

import (
	"github.com/warren-community/warren/contenttype/text"
)

var DefaultContentType = new(text.Plain)

var Registry = map[string]ContentType{
	"text/plain":    new(text.Plain),
	"text/markdown": new(text.Markdown),
}

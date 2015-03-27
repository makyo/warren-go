package user

import (
	"github.com/warren-community/warren/contenttype/text"
)

type Profile struct{}

// Since users are managed through markdown, they are a safe content type.
func (c *Profile) Safe() bool {
	return true
}

// Render the profile using markdown
// TODO Users may need additional fields in the future.
func (c *Profile) RenderDisplayContent(content string) (string, error) {
	ct := new(text.Markdown)
	return ct.RenderDisplayContent(content)
}

// Simply return the markdown content.
func (c *Profile) RenderIndexContent(content string) (string, error) {
	return content, nil
}

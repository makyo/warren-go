package text

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type Markdown struct{}

func (c *Markdown) Safe() bool {
	return true
}

func (c *Markdown) RenderDisplayContent(content string) (string, error) {
	rendered := blackfriday.MarkdownCommon([]byte(content))
	safe := bluemonday.UGCPolicy().SanitizeBytes(rendered)
	return string(safe), nil
}

func (c *Markdown) RenderIndexContent(content string) (string, error) {
	return content, nil
}

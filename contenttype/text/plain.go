package text

import (
	"fmt"
	"html/template"
	"regexp"
)

type Plain struct{}

func (c *Plain) Safe() bool {
	return true
}

func (c *Plain) RenderDisplayContent(content string) (string, error) {
	content = template.HTMLEscapeString(content)
	paraRe, err := regexp.Compile("\r\n\r\n")
	if err != nil {
		return "", err
	}
	breakRe, err := regexp.Compile("\r\n")
	if err != nil {
		return "", err
	}
	content = string(paraRe.ReplaceAll([]byte(content), []byte("</p><p>")))
	content = string(breakRe.ReplaceAll([]byte(content), []byte("<br />")))
	return fmt.Sprintf("<p>%s</p>", content), nil
}

func (c *Plain) RenderIndexContent(content string) (string, error) {
	return content, nil
}

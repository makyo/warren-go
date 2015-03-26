package text

type Raw struct{}

func (c *Raw) Safe() bool {
	return false
}

func (c *Raw) RenderDisplayContent(content string) (string, error) {
	return content, nil
}

func (c *Raw) RenderIndexContent(content string) (string, error) {
	return content, nil
}

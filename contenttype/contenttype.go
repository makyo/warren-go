package contenttype

type ContentType interface {
	RenderDisplayContent(content string) (string, error)
	RenderIndexContent(content string) (string, error)
	Safe() bool
}

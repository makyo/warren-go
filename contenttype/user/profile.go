package user

import (
	"bytes"
	"fmt"
	"html/template"

	"gopkg.in/mgo.v2/bson"

	"github.com/warren-community/warren/contenttype/text"
	"github.com/warren-community/warren/utils"
)

type Profile struct {
	ProfileText string
	Pronouns    string
	Website     string
}

// Create a new Profile object from a mongo result.
func NewProfile(in bson.M) Profile {
	fields := []string{
		"profiletext",
		"pronouns",
		"website",
	}
	for _, field := range fields {
		_, ok := in[field]
		if !ok {
			in[field] = ""
		}
	}
	return Profile{
		ProfileText: in["profiletext"].(string),
		Pronouns:    in["pronouns"].(string),
		Website:     in["website"].(string),
	}
}

// Since users are managed through markdown, they are a safe content type.
func (c *Profile) Safe() bool {
	return true
}

// Render the profile using markdown
// TODO Users may need additional fields in the future.
func (c *Profile) RenderDisplayContent(content interface{}) (string, error) {
	profile := content.(Profile)
	profileText := template.HTML(text.RenderMarkdown(profile.ProfileText))
	buf := new(bytes.Buffer)
	td, err := utils.GetTemplateDir()
	if err != nil {
		return "", err
	}
	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("%s/templates/contenttype/user/profile.tmpl", td)))
	tmpl.Execute(buf, map[string]interface{}{
		"ProfileText": profileText,
		"Website":     profile.Website,
		"Pronouns":    profile.Pronouns,
	})
	return buf.String(), nil
}

// Simply return the profile text content.
func (c *Profile) RenderIndexContent(content interface{}) (string, error) {
	return (content.(Profile)).ProfileText, nil
}

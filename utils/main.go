package utils

import (
	"os"
	"regexp"
)

func GetTemplateDir() (string, error) {
	re := regexp.MustCompile("warren/(.*)$")
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	wd = re.ReplaceAllString(wd, "warren")
	return wd, nil
}

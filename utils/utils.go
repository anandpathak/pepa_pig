package utils

import (
	"regexp"
	"strings"
)

func ConvertTolowerCaseWithoutSpace(str string) string {
	space := regexp.MustCompile(`\s+`)
	s := space.ReplaceAllString(strings.TrimSpace(str), "_")
	return strings.ToLower(s)

}

package flagbind

import (
	"regexp"
	"strings"
)

var (
	matchFirstCapitalRgx = regexp.MustCompile(`(.)([A-Z][a-z]+)`)
	matchAllCapitalRgx   = regexp.MustCompile(`([a-z0-9])([A-Z])`)
)

func toArgName(str string) string {
	result := matchFirstCapitalRgx.ReplaceAllString(str, "${1}-${2}")
	result = matchAllCapitalRgx.ReplaceAllString(result, "${1}-${2}")
	return strings.ToLower(result)
}

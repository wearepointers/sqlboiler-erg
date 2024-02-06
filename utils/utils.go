package utils

import (
	"strings"
	"unicode"

	"github.com/gedex/inflector"
)

func toSnakeCase(s string) string {
	var result []rune

	s = strings.ReplaceAll(s, "ID", "id")
	s = strings.ReplaceAll(s, "URL", "url")
	s = strings.ReplaceAll(s, "JSON", "json")
	s = strings.ReplaceAll(s, "API", "api")

	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}

	return string(result)
}

func toCamelCase(s string) string {
	var result []rune

	for i, r := range s {
		if i == 0 {
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}

	return string(result)
}

func pluralize(s string) string {
	return inflector.Pluralize(s)
}

func singularize(s string) string {
	return inflector.Singularize(s)
}

func titleize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
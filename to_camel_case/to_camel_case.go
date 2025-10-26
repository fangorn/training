package to_camel_case

import (
	"strings"
	"unicode"
)

// https://www.codewars.com/kata/517abf86da9663f1d2000003/train/go
func ToCamelCase(s string) string {
	dashed := strings.Replace(s, "_", "-", -1)
	parts := strings.Split(dashed, "-")

	var result string
	for i, p := range parts {
		runes := []rune(p)
		if i != 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		result += string(runes)
	}
	return result
}

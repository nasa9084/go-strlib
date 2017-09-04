package strlib

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

// SnakeCase converts camel case to snake case.
func SnakeCase(s string) string {
	return snakeCase(s, []rune{}, false)
}

func snakeCase(s string, rs []rune, wasLower bool) string {
	if len(s) == 0 {
		return string(rs)
	}

	r, size := utf8.DecodeRuneInString(s)
	s = s[size:]

	if unicode.IsUpper(r) {
		if wasLower {
			rs = append(rs, '_')
		}
		wasLower = false
	} else { // if IsLower
		wasLower = true
	}
	rs = append(rs, unicode.ToLower(r))

	return snakeCase(s, rs, wasLower)
}

// UpperCamelCase converts snake case to upper camel case
func UpperCamelCase(s string) string {
	ss := strings.Split(s, "_")
	if len(ss) == 1 {
		return strings.Title(s)
	}
	buf := bytes.Buffer{}
	for _, p := range ss {
		buf.WriteString(strings.Title(p))
	}
	return buf.String()
}

// LowerCamelCase converts snake case to lower camel case
func LowerCamelCase(s string) string {
	ss := strings.Split(s, "_")
	if len(ss) == 1 {
		return strings.ToLower(string(s[0])) + s[1:]
	}
	buf := bytes.Buffer{}
	for i, p := range ss {
		if i == 0 {
			buf.WriteString(strings.ToLower(string(p[0])) + p[1:])
			continue
		}
		buf.WriteString(strings.ToUpper(string(p[0])) + strings.ToLower(p[1:]))
	}
	return buf.String()
}

// Capitalize returns a copy of the string with its first character capitalized and the rest lowercased
func Capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

// IsUpper returns true if all cased characters in the string are uppercase
func IsUpper(s string) bool {
	if len(s) < 1 {
		return false
	}
	return strings.ToUpper(s) == s
}

// IsLower returns true if all cased characters in the string are lowercase
func IsLower(s string) bool {
	if len(s) < 1 {
		return false
	}
	return strings.ToLower(s) == s
}

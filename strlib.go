package strlib

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

// SnakeCase converts camel case to snake case.
func SnakeCase(s string) string {
	wasLower := false
	rs := []rune{}

	for len(s) > 0 {
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
	}
	return string(rs)
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

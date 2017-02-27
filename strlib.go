package strlib

import (
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

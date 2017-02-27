package strlib_test

import (
	"testing"

	strlib "github.com/nasa9084/go-strlib"
)

func TestSnakeCase(t *testing.T) {
	candidates := []struct {
		camel string
		snake string
	}{
		{"lowerCamelCaseString", "lower_camel_case_string"},
		{"UpperCamelCaseString", "upper_camel_case_string"},
		{"snake_case_string", "snake_case_string"},
	}

	for _, c := range candidates {
		t.Logf("Do: %s\n", c.camel)
		t.Logf("  Expected: %s\n", c.snake)
		t.Logf("  Got: %s\n", strlib.SnakeCase(c.camel))
		if strlib.SnakeCase(c.camel) != c.snake {
			t.Errorf("%s != %s\n", strlib.SnakeCase(c.camel), c.snake)
			return
		}
	}
}

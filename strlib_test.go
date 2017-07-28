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

func TestUpperCamelCase(t *testing.T) {
	candidates := []struct {
		snake string
		camel string
	}{
		{"snake_case_string", "SnakeCaseString"},
		{"lowerCamelCaseString", "LowerCamelCaseString"},
		{"UpperCamelCaseString", "UpperCamelCaseString"},
	}
	for _, c := range candidates {
		t.Logf("Do: %s\n", c.camel)
		t.Logf("  Expected: %s\n", c.snake)
		t.Logf("  Got: %s\n", strlib.UpperCamelCase(c.snake))
		if strlib.UpperCamelCase(c.snake) != c.camel {
			t.Errorf("%s != %s\n", strlib.UpperCamelCase(c.snake), c.camel)
		}
	}

}

func TestLowerCamelCase(t *testing.T) {
	candidates := []struct {
		snake string
		camel string
	}{
		{"snake_case_string", "snakeCaseString"},
		{"lowerCamelCaseString", "lowerCamelCaseString"},
		{"UpperCamelCaseString", "upperCamelCaseString"},
	}
	for _, c := range candidates {
		t.Logf("Do: %s\n", c.camel)
		t.Logf("  Expected: %s\n", c.snake)
		t.Logf("  Got: %s\n", strlib.LowerCamelCase(c.snake))
		if strlib.LowerCamelCase(c.snake) != c.camel {
			t.Errorf("%s != %s\n", strlib.LowerCamelCase(c.snake), c.camel)
		}
	}

}

func TestCapitalize(t *testing.T) {
	candidates := []struct {
		in  string
		out string
	}{
		{"lower", "Lower"},
		{"Upper", "Upper"},
		{"ALLUPPER", "Allupper"},
		{"lower with spaces", "Lower with spaces"},
		{"Upper with spaces", "Upper with spaces"},
		{"Upper Title", "Upper title"},
	}
	for _, c := range candidates {
		got := strlib.Capitalize(c.in)
		t.Logf("Do: %s\n", c.in)
		t.Logf("  Expected: %s\n", c.out)
		t.Logf("  Got: %s\n", got)
		if got != c.out {
			t.Errorf("%s != %s\n", got, c.out)
		}
	}
}

package goescapemysql

import (
	"testing"
	"unicode/utf8"
)

type args struct {
	value string
}

func getWant(s string) string {
	switch s[0] {
	case '\n':
		return `\n`
	case '\r':
		return `\r`
	case 0:
		return `\0`
	case '\\':
		return `\\`
	case '\'':
		return `\'`
	case '"':
		return `\"`
	case '\032':
		return `\Z`
	case '_':
		return `\_`
	case '%':
		return `\%`
	}
	return s
}

// TestRealEscapeString tests RealEscapeString.
func TestRealEscapeString(t *testing.T) {
	for r := rune(0); r < utf8.MaxRune; r++ {
		if !utf8.ValidRune(r) {
			continue
		}
		given := string(r)
		want := getWant(given)
		got := RealEscapeString(given)
		if want != got {
			t.Errorf("RealEscapeString() = %v, want %v", got, want)
		}
	}
}

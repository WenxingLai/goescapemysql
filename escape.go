package goescapemysql

import (
	"strings"
)

// RealEscapeString escapes a string as mysql-server does.
func RealEscapeString(value string) string {
	var sb strings.Builder
	// Source: #789 escape_string_for_mysql https://github.com/mysql/mysql-server/blob/5.7/mysys/charset.c
	// Adding: % and _
	for _, v := range value {
		switch v {
		case '\n':
			sb.WriteByte('\\')
			sb.WriteByte('n')
		case '\r':
			sb.WriteByte('\\')
			sb.WriteByte('r')
		case 0:
			sb.WriteByte('\\')
			sb.WriteByte('0')
		case '\\':
			sb.WriteByte('\\')
			sb.WriteByte('\\')
		case '\'':
			sb.WriteByte('\\')
			sb.WriteByte('\'')
		case '"':
			sb.WriteByte('\\')
			sb.WriteByte('"')
		case '\032':
			sb.WriteByte('\\')
			sb.WriteByte('Z')
		case '_':
			sb.WriteByte('\\')
			sb.WriteByte('_')
		case '%':
			sb.WriteByte('\\')
			sb.WriteByte('%')
		default:
			sb.WriteRune(v)
		}
	}
	return sb.String()
}

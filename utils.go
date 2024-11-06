package goblock

import "strings"

func splitPath(path string) []string {
	if path == "/" || len(path) == 0 {
		return []string{"/"}
	}

	return strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})
}

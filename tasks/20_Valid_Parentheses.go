package main

import (
	"strings"
)

func isValid(s string) bool {
	for s != "" {
		old_s := s
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		if old_s == s {
			return false
		}
	}
	return true
}

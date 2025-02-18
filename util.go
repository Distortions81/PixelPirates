package main

import "strings"

func removeNewlines(input string) string {
	input = strings.ReplaceAll(input, "\n", "")
	return strings.ReplaceAll(input, "\r", "")
}

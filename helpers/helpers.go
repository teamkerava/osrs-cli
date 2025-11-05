package helpers

import (
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// formats large numbers with commas for better readability
func FormatNumber(n int) string {
	str := strconv.Itoa(n)
	length := len(str)

	if length <= 3 {
		return str
	}

	var result strings.Builder
	for i, digit := range str {
		if i > 0 && (length-i)%3 == 0 {
			result.WriteString(",")
		}
		result.WriteRune(digit)
	}

	return result.String()
}

// formats a name by replacing underscores with spaces and capitalizing each word
func FormatName(name string) string {
	name = strings.ReplaceAll(name, "_", " ")
	c := cases.Title(language.English)
	return c.String(name)
}

func MaxLen(current int, value string) int {
	if len(value) > current {
		return len(value)
	}
	return current
}

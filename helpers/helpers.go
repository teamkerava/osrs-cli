package helpers

import (
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// FormatNumber formats large numbers with commas for better readability
func FormatNumber(n int) string {
	str := strconv.Itoa(n)
	length := len(str)

	if length <= 3 {
		return str
	}

	// Insert commas every 3 digits from the right
	var result strings.Builder
	for i, digit := range str {
		if i > 0 && (length-i)%3 == 0 {
			result.WriteString(",")
		}
		result.WriteRune(digit)
	}

	return result.String()
}

// Format Name to have Uppercase Titles
func FormatName(name string) string {
	name = strings.ReplaceAll(name, "_", " ")
	c := cases.Title(language.English)
	return c.String(name)
}

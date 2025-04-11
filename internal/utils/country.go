package utils

import (
	"net/url"
	"strings"
	"unicode"
)

// capitalizeWord capitalizes the first rune of a word and lowercases the rest
func capitalizeWord(word string) string {
	if word == "" {
		return ""
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

// NormalizeCountryName decodes a URL-encoded string and returns a title-cased country name.
func NormalizeCountryName(raw string) (string, error) {
	decoded, err := url.QueryUnescape(raw)
	if err != nil {
		return "", err
	}

	words := strings.Fields(decoded)
	for i, word := range words {
		words[i] = capitalizeWord(word)
	}

	return strings.Join(words, " "), nil
}

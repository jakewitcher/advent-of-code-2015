package rules

import (
	"strings"
)

type Rule func(string) bool

func HasNoProhibitedSubstring(text string) bool {
	prohibitedSubstrings := []string{"ab", "cd", "pq", "xy"}

	for _, subString := range prohibitedSubstrings {
		if strings.Contains(text, subString) {
			return false
		}
	}

	return true
}

func HasAdjacentDuplicateLetters(text string) bool {
	prev := rune(text[0])
	for i, r := range text {
		if prev == r && i != 0 {
			return true
		}

		prev = r
	}

	return false
}

func HasThreeVowels(text string) bool {
	vowels := "aeiou"
	var vowelsCount int

	for _, r := range text {
		if strings.Contains(vowels, string(r)) {
			vowelsCount++
		}
	}

	return vowelsCount >= 3
}

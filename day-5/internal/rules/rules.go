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
	for _, r := range text[1:] {
		if prev == r {
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

func HasPairOfLettersAppearingTwice(text string) bool {
	seenPairs := make(map[string]int)

	var prevRune rune
	var pairStartPosition int

	for i, r := range text {
		if i == 0 {
			prevRune = r
			continue
		}

		pair := string(prevRune) + string(r)
		pairEndPosition, ok := seenPairs[pair]

		if ok && pairStartPosition != pairEndPosition {
			return true
		}

		if !ok {
			seenPairs[pair] = i
		}

		prevRune = r
		pairStartPosition = i
	}

	return false
}

func HasRepeatingLetterWithOneLetterBetween(text string) bool {
	prevEven := rune(text[0])
	prevOdd := rune(text[1])

	for i, r := range text[2:] {
		if i%2 == 0 {
			if r == prevEven {
				return true
			}

			prevEven = r
		} else {
			if r == prevOdd {
				return true
			}

			prevOdd = r
		}
	}

	return false
}

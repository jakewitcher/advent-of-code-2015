package test

import (
	"day-5/internal/rules"
	"day-5/internal/text"
	"log"
	"testing"
)

var partOneRules = []rules.Rule{
	rules.HasNoProhibitedSubstring,
	rules.HasAdjacentDuplicateLetters,
	rules.HasThreeVowels,
}

func TestEvaluateTextDemeanor(t *testing.T) {
	for _, test := range singleTextTestCases {
		if demeanor := text.EvaluateDemeanor(test.text, partOneRules...); demeanor != test.demeanor {
			log.Fatalf("for text %s expected demeanor: %s, actual demeanor: %s", test.text, test.demeanor, demeanor)
		}
	}
}

func TestEvaluateTextDemeanorAll(t *testing.T) {
	for _, test := range allTextTestCases {
		if niceCount := text.EvaluateDemeanorAll(test.texts, partOneRules...); niceCount != test.niceCount {
			log.Fatalf("expected nice count: %d, actual nice count: %d", test.niceCount, niceCount)
		}
	}
}

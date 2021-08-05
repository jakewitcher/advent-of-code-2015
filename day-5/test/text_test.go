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

var partTwoRules = []rules.Rule{
	rules.HasPairOfLettersAppearingTwice,
	rules.HasRepeatingLetterWithOneLetterBetween,
}

func TestEvaluateTextDemeanorUsingPartOneRules(t *testing.T) {
	for _, test := range singleTextPartOneRulesTestCases {
		if demeanor := text.EvaluateDemeanor(test.text, partOneRules...); demeanor != test.demeanor {
			log.Fatalf("for text %s expected demeanor: %s, actual demeanor: %s", test.text, test.demeanor, demeanor)
		}
	}
}

func TestEvaluateTextDemeanorAllUsingPartOneRules(t *testing.T) {
	for _, test := range allTextPartOneRulesTestCases {
		if niceCount := text.EvaluateDemeanorAll(test.texts, partOneRules...); niceCount != test.niceCount {
			log.Fatalf("expected nice count: %d, actual nice count: %d", test.niceCount, niceCount)
		}
	}
}

func TestEvaluateTextDemeanorUsingPartTwoRules(t *testing.T) {
	for _, test := range singleTextPartTwoRulesTestCases {
		if demeanor := text.EvaluateDemeanor(test.text, partTwoRules...); demeanor != test.demeanor {
			log.Fatalf("for text %s expected demeanor: %s, actual demeanor: %s", test.text, test.demeanor, demeanor)
		}
	}
}

func TestEvaluateTextDemeanorAllUsingPartTwoRules(t *testing.T) {
	for _, test := range allTextPartTwoRulesTestCases {
		if niceCount := text.EvaluateDemeanorAll(test.texts, partTwoRules...); niceCount != test.niceCount {
			log.Fatalf("expected nice count: %d, actual nice count: %d", test.niceCount, niceCount)
		}
	}
}

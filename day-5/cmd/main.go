package main

import (
	"day-5/internal/input"
	"day-5/internal/rules"
	"day-5/internal/text"
	"log"
)

func main() {
	texts, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	PartOne(texts)
	PartTwo(texts)
}

func PartOne(texts []string) {
	partOneRules := []rules.Rule{
		rules.HasNoProhibitedSubstring,
		rules.HasAdjacentDuplicateLetters,
		rules.HasThreeVowels,
	}
	niceCount := text.EvaluateDemeanorAll(texts, partOneRules...)
	log.Printf("nice count: %d", niceCount)
}

func PartTwo(texts []string) {
	partTwoRules := []rules.Rule{
		rules.HasPairOfLettersAppearingTwice,
		rules.HasRepeatingLetterWithOneLetterBetween,
	}
	niceCount := text.EvaluateDemeanorAll(texts, partTwoRules...)
	log.Printf("nice count: %d", niceCount)
}

package text

import (
	"day-5/internal/rules"
	"fmt"
)

const (
	Naughty = "naughty"
	Nice    = "nice"
)

func EvaluateDemeanorAll(texts []string, rules ...rules.Rule) int {
	var niceCount int

	for i, text := range texts {
		if demeanor := EvaluateDemeanor(text, rules...); demeanor == Nice {
			fmt.Println(i + 1)
			niceCount++
		}
	}

	return niceCount
}

func EvaluateDemeanor(text string, rules ...rules.Rule) string {
	isNice := true

	for _, rule := range rules {
		isNice = isNice && rule(text)
	}

	if isNice {
		return Nice
	}
	return Naughty
}

package text

import (
	"day-5/internal/rules"
)

const (
	Naughty = "naughty"
	Nice    = "nice"
)

func EvaluateDemeanorAll(texts []string, rules ...rules.Rule) int {
	var niceCount int

	for _, text := range texts {
		if demeanor := EvaluateDemeanor(text, rules...); demeanor == Nice {
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

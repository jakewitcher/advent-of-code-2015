package test

import "day-5/internal/text"

var singleTextPartOneRulesTestCases = []struct {
	text     string
	demeanor string
}{
	{
		text:     "ugknbfddgicrmopn",
		demeanor: text.Nice,
	},
	{
		text:     "aaa",
		demeanor: text.Nice,
	},
	{
		text:     "jchzalrnumimnmhp",
		demeanor: text.Naughty,
	},
	{
		text:     "haegwjzuvuyypxyu",
		demeanor: text.Naughty,
	},
	{
		text:     "dvszwmarrgswjxmb",
		demeanor: text.Naughty,
	},
}

var allTextPartOneRulesTestCases = []struct {
	texts     []string
	niceCount int
}{
	{
		texts: []string{
			"ugknbfddgicrmopn",
			"aaa",
			"jchzalrnumimnmhp",
			"haegwjzuvuyypxyu",
			"dvszwmarrgswjxmb",
		},
		niceCount: 2,
	},
}

var singleTextPartTwoRulesTestCases = []struct {
	text     string
	demeanor string
}{
	{
		text:     "qjhvhtzxzqqjkmpb",
		demeanor: text.Nice,
	},
	{
		text:     "xxyxx",
		demeanor: text.Nice,
	},
	{
		text:     "uurcxstgmygtbstg",
		demeanor: text.Naughty,
	},
	{
		text:     "ieodomkazucvgmuy",
		demeanor: text.Naughty,
	},
	{
		text:     "aaaba",
		demeanor: text.Naughty,
	},
	{
		text:     "sknufchjdvccccta",
		demeanor: text.Nice,
	},
}

var allTextPartTwoRulesTestCases = []struct {
	texts     []string
	niceCount int
}{
	{
		texts: []string{
			"qjhvhtzxzqqjkmpb",
			"xxyxx",
			"uurcxstgmygtbstg",
			"ieodomkazucvgmuy",
			"aaaba",
		},
		niceCount: 2,
	},
}

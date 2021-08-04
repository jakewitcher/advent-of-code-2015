package test

import "day-5/internal/text"

var singleTextTestCases = []struct {
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

var allTextTestCases = []struct {
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

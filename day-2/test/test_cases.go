package test

import "day-2/internal/domain"

var calculateWrappingPaperRequiredTestCases = []struct {
	dimensions []string
	squareFeet int
}{
	{
		dimensions: []string{"2x3x4", "1x1x10"},
		squareFeet: 101,
	},
}

var calculateRibbonRequiredTestCases = []struct {
	dimensions []string
	feet       int
}{
	{
		dimensions: []string{"2x3x4", "1x1x10"},
		feet:       48,
	},
}

var parseDimensionsTestCases = []struct {
	dimensions string
	present    *domain.Present
}{
	{
		dimensions: "2x3x4",
		present:    domain.NewPresent(2, 3, 4),
	},
	{
		dimensions: "1x1x10",
		present:    domain.NewPresent(1, 1, 10),
	},
}

var wrappingPaperSqFtTestCases = []struct {
	present *domain.Present
	sqFt    int
}{
	{
		present: domain.NewPresent(2, 3, 4),
		sqFt:    58,
	},
	{
		present: domain.NewPresent(1, 1, 10),
		sqFt:    43,
	},
}

var ribbonFtTestCases = []struct {
	present *domain.Present
	ft      int
}{
	{
		present: domain.NewPresent(2, 3, 4),
		ft:      34,
	},
	{
		present: domain.NewPresent(1, 1, 10),
		ft:      14,
	},
	{
		present: domain.NewPresent(7, 1, 10),
		ft:      86,
	},
}

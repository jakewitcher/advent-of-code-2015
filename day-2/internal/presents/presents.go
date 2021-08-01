package presents

import (
	"day-2/internal/domain"
	"strconv"
	"strings"
)

func CalculateWrappingPaperRequired(dimensions []string) (int, error) {
	var sqFt int

	for _, d := range dimensions {
		present, err := ParseDimensions(d)
		if err != nil {
			return 0, err
		}

		sqFt += present.WrappingPaperSqFt()
	}

	return sqFt, nil
}

func CalculateRibbonRequired(dimensions []string) (int, error) {
	var ft int

	for _, d := range dimensions {
		present, err := ParseDimensions(d)
		if err != nil {
			return 0, err
		}

		ft += present.RibbonFt()
	}

	return ft, nil
}

func ParseDimensions(dimensions string) (*domain.Present, error) {
	values := strings.Split(dimensions, "x")

	length, err := strconv.Atoi(values[0])
	if err != nil {
		return nil, err
	}

	width, err := strconv.Atoi(values[1])
	if err != nil {
		return nil, err
	}

	height, err := strconv.Atoi(values[2])
	if err != nil {
		return nil, err
	}

	present := domain.NewPresent(length, width, height)

	return present, nil
}

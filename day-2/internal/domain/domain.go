package domain

import "math"

type Present struct {
	length int
	width  int
	height int
}

func (p Present) WrappingPaperSqFt() int {
	a := p.length * p.width
	b := p.length * p.height
	c := p.width * p.height

	return a*2 + b*2 + c*2 + smallest(a, b, c)
}

func (p Present) RibbonFt() int {
	a, b := twoSmallest(p.length, p.width, p.height)
	c := p.length * p.width * p.height

	return a*2 + b*2 + c
}

func smallest(nums ...int) int {
	min := nums[0]

	for _, n := range nums {
		if n < min {
			min = n
		}
	}

	return min
}

func twoSmallest(nums ...int) (int, int) {
	var min1, min2 = math.MaxInt32, math.MaxInt32

	for _, n := range nums {
		if n < min1 {
			min1, min2 = n, min1
		} else if n < min2 {
			min2 = n
		}
	}

	return min1, min2
}

func NewPresent(length int, width int, height int) *Present {
	return &Present{
		length,
		width,
		height,
	}
}

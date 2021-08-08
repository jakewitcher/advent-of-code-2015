package test

import (
	d "day-6/internal/direction"
	"testing"
)

func TestParseDirection(t *testing.T) {
	for _, test := range parseDirectionTestCases {
		if direction := d.Parse(test.rawDirection); direction != test.direction {
			t.Fatalf("expected direction: %v, actual direction %v", test.direction, direction)
		}
	}
}

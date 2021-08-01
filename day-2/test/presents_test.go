package test

import (
	"day-2/internal/input"
	"day-2/internal/presents"
	"testing"
)

func TestCalculateWrappingPaperRequired(t *testing.T) {
	for _, test := range calculateWrappingPaperRequiredTestCases {
		if squareFeet, _ := presents.CalculateWrappingPaperRequired(test.dimensions); squareFeet != test.squareFeet {
			t.Fatalf("expected sq feet: %d, actual sq feet: %d", test.squareFeet, squareFeet)
		}
	}
}

func TestCalculateRibbonRequired(t *testing.T) {
	for _, test := range calculateRibbonRequiredTestCases {
		if feet, _ := presents.CalculateRibbonRequired(test.dimensions); feet != test.feet {
			t.Fatalf("expected ft: %d, actual ft: %d", test.feet, feet)
		}
	}
}

func TestParseDimensions(t *testing.T) {
	for _, test := range parseDimensionsTestCases {
		if present, _ := presents.ParseDimensions(test.dimensions); *present != *test.present {
			t.Fatalf("expected present: %v, actual present: %v", *test.present, *present)
		}
	}
}

func BenchmarkCalculateWrappingPaperRequired(b *testing.B) {
	dimensions, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = presents.CalculateWrappingPaperRequired(dimensions)
	}
}

func BenchmarkCalculateRibbonRequired(b *testing.B) {
	dimensions, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = presents.CalculateRibbonRequired(dimensions)
	}
}

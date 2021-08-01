package test

import "testing"

func TestWrappingPaperSqFt(t *testing.T) {
	for _, test := range wrappingPaperSqFtTestCases {
		if sqFt := test.present.WrappingPaperSqFt(); sqFt != test.sqFt {
			t.Fatalf("expected sq ft: %d, actual sq ft: %d", test.sqFt, sqFt)
		}
	}
}

func TestRibbonFt(t *testing.T) {
	for _, test := range ribbonFtTestCases {
		if ft := test.present.RibbonFt(); ft != test.ft {
			t.Fatalf("expected ft: %d, actual ft: %d", test.ft, ft)
		}
	}
}

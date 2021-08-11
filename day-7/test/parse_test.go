package test

import (
	"day-7/parse"
	"testing"
)

func TestParseOneInstruction(t *testing.T) {
	for _, test := range parseInstructionTestCases {
		if actualInstruction, _ := parse.One(test.rawInstruction); actualInstruction != test.expectedInstruction {
			t.Fatalf("expected instruction: %v, actual instruction: %v", test.expectedInstruction, actualInstruction)
		}
	}
}

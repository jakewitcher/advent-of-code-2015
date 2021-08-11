package test

import (
	"day-7/parse"
	"testing"
)

func TestApplyInstruction(t *testing.T) {
	for _, test := range applyInstructionTestCases {
		if actualWire := test.instruction.Apply(); actualWire != test.expectedWire {
			t.Fatalf("expected wire: %v, actual wire: %v", test.expectedWire, actualWire)
		}
	}
}

func TestParseOneInstruction(t *testing.T) {
	for _, test := range parseInstructionTestCases {
		if actualInstruction, _ := parse.One(test.rawInstruction); actualInstruction != test.expectedInstruction {
			t.Fatalf("expected instruction: %v, actual instruction: %v", test.expectedInstruction, actualInstruction)
		}
	}
}
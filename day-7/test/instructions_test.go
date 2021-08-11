package test

import (
	"testing"
)

func TestApplyInstruction(t *testing.T) {
	for _, test := range applyInstructionTestCases {
		if actualWire := test.instruction.Apply(); actualWire != test.expectedWire {
			t.Fatalf("expected wire: %v, actual wire: %v", test.expectedWire, actualWire)
		}
	}
}

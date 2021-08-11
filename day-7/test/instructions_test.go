package test

import (
	"testing"
)

func TestApplyInstructions(t *testing.T) {
	for _, test := range applyInstructionTestCases {
		if actualWire := test.instruction.Apply(); actualWire != test.expectedWire {
			t.Fatalf("expected wire: %v, actual wire: %v", test.expectedWire, actualWire)
		}
	}
}


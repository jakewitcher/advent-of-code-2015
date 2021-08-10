package test

import (
	"day-7/internal/wires"
	"testing"
)

func TestApplyAndOrSignal(t *testing.T) {
	for _, test := range andOrTestCases {
		actualWire := wires.ApplyAndOrSignal(test.leftInputWire, test.rightInputWire, test.gate, test.inputIdentifier)
		if actualWire != test.expectedWire {
			t.Fatalf("expected wire: %v, actual wire: %v", test.expectedWire, actualWire)
		}
	}
}

func TestApplyLShiftRShiftSignal(t *testing.T) {
	for _, test := range lShiftRShiftTestCases {
		actualWire := wires.ApplyLShiftRShiftSignal(test.inputWire, test.gate, test.shift, test.inputIdentifier)
		if actualWire != test.expectedWire {
			t.Fatalf("expected wire: %v, actual wire: %v", test.expectedWire, actualWire)
		}
	}
}

func TestApplyNotSignal(t *testing.T) {
	for _, test := range notTestCases {
		actualWire := wires.ApplyNotSignal(test.inputWire, test.inputIdentifier)
		if actualWire != test.expectedWire {
			t.Fatalf("expected wire: %v, actual wire: %v", test.expectedWire, actualWire)
		}
	}
}

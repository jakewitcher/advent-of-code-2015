package wires

import "day-7/internal/domain"

func ApplyAndOrSignal(leftWire domain.Wire, rightWire domain.Wire, gate domain.Gate, identifier domain.Identifier) domain.Wire {
	var newSignal domain.Signal

	switch gate {
	case domain.And:
		newSignal = leftWire.Signal & rightWire.Signal
	case domain.Or:
		newSignal = leftWire.Signal | rightWire.Signal
	}

	return domain.NewWire(identifier, newSignal)
}

func ApplyLShiftRShiftSignal(inputWire domain.Wire, gate domain.Gate, shift domain.Shift, identifier domain.Identifier) domain.Wire {
	var newSignal domain.Signal

	switch gate {
	case domain.LShift:
		newSignal = inputWire.Signal << shift
	case domain.RShift:
		newSignal = inputWire.Signal >> shift
	}

	return domain.NewWire(identifier, newSignal)
}

func ApplyNotSignal(inputWire domain.Wire, identifier domain.Identifier) domain.Wire {
	newSignal := ^inputWire.Signal
	return domain.NewWire(identifier, newSignal)
}

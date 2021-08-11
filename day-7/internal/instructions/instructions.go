package instructions

import "day-7/internal/wires"

type InstructionApplier interface {
	Apply() wires.Wire
}

type SignalAssignmentInstruction struct {
	producer   wires.SignalProducer
	identifier wires.Identifier
}

func (i SignalAssignmentInstruction) Apply() wires.Wire {
	return wires.NewWire(i.identifier, i.producer.ProduceSignal())
}

func NewSignalAssignmentInstruction(producer wires.SignalProducer, identifier wires.Identifier) InstructionApplier {
	return SignalAssignmentInstruction{
		producer:   producer,
		identifier: identifier,
	}
}

type AndInstruction struct {
	left       wires.SignalProducer
	right      wires.SignalProducer
	identifier wires.Identifier
}

func (i AndInstruction) Apply() wires.Wire {
	return wires.NewWire(i.identifier, i.left.ProduceSignal()&i.right.ProduceSignal())
}

func NewAndInstruction(left wires.SignalProducer, right wires.SignalProducer, identifier wires.Identifier) InstructionApplier {
	return AndInstruction{
		left:       left,
		right:      right,
		identifier: identifier,
	}
}

type OrInstruction struct {
	left       wires.SignalProducer
	right      wires.SignalProducer
	identifier wires.Identifier
}

func (i OrInstruction) Apply() wires.Wire {
	return wires.NewWire(i.identifier, i.left.ProduceSignal()|i.right.ProduceSignal())
}

func NewOrInstruction(left wires.SignalProducer, right wires.SignalProducer, identifier wires.Identifier) InstructionApplier {
	return OrInstruction{
		left:       left,
		right:      right,
		identifier: identifier,
	}
}

type LShiftInstruction struct {
	producer   wires.SignalProducer
	identifier wires.Identifier
	shift      wires.Shift
}

func (i LShiftInstruction) Apply() wires.Wire {
	return wires.NewWire(i.identifier, i.producer.ProduceSignal()<<i.shift)
}

func NewLShiftInstruction(producer wires.SignalProducer, shift wires.Shift, identifier wires.Identifier) InstructionApplier {
	return LShiftInstruction{
		producer:   producer,
		shift:      shift,
		identifier: identifier,
	}
}

type RShiftInstruction struct {
	producer   wires.SignalProducer
	identifier wires.Identifier
	shift      wires.Shift
}

func (i RShiftInstruction) Apply() wires.Wire {
	return wires.NewWire(i.identifier, i.producer.ProduceSignal()>>i.shift)
}

func NewRShiftInstruction(producer wires.SignalProducer, shift wires.Shift, identifier wires.Identifier) InstructionApplier {
	return RShiftInstruction{
		producer:   producer,
		shift:      shift,
		identifier: identifier,
	}
}

type NotInstruction struct {
	producer   wires.SignalProducer
	identifier wires.Identifier
}

func (i NotInstruction) Apply() wires.Wire {
	return wires.NewWire(i.identifier, ^i.producer.ProduceSignal())
}

func NewNotInstruction(producer wires.SignalProducer, identifier wires.Identifier) InstructionApplier {
	return NotInstruction{
		producer:   producer,
		identifier: identifier,
	}
}

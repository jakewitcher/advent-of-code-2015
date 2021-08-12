package instructions

import "day-7/internal/wires"

type InstructionApplier interface {
	Apply() wires.Wire
}

type WireIdentifier interface {
	Identify() wires.Identifier
}

type Instruction interface {
	InstructionApplier
	WireIdentifier
}

type SignalAssignmentInstruction struct {
	producer   wires.SignalProducer
	identifier wires.Identifier
}

func (i SignalAssignmentInstruction) Apply() wires.Wire {
	return wires.NewWire(i.identifier, i.producer.ProduceSignal())
}

func (i SignalAssignmentInstruction) Identify() wires.Identifier {
	return i.identifier
}

func NewSignalAssignmentInstruction(producer wires.SignalProducer, identifier wires.Identifier) Instruction {
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

func (i AndInstruction) Identify() wires.Identifier {
	return i.identifier
}

func NewAndInstruction(left wires.SignalProducer, right wires.SignalProducer, identifier wires.Identifier) Instruction {
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

func (i OrInstruction) Identify() wires.Identifier {
	return i.identifier
}

func NewOrInstruction(left wires.SignalProducer, right wires.SignalProducer, identifier wires.Identifier) Instruction {
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

func (i LShiftInstruction) Identify() wires.Identifier {
	return i.identifier
}

func NewLShiftInstruction(producer wires.SignalProducer, shift wires.Shift, identifier wires.Identifier) Instruction {
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

func (i RShiftInstruction) Identify() wires.Identifier {
	return i.identifier
}

func NewRShiftInstruction(producer wires.SignalProducer, shift wires.Shift, identifier wires.Identifier) Instruction {
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

func (i NotInstruction) Identify() wires.Identifier {
	return i.identifier
}

func NewNotInstruction(producer wires.SignalProducer, identifier wires.Identifier) Instruction {
	return NotInstruction{
		producer:   producer,
		identifier: identifier,
	}
}

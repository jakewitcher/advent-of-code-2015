package instructions

import "day-7/internal/wires"

type Instruction interface {
	Apply() wires.Wire
}

type AndInstruction struct {
	left       wires.Wire
	right      wires.Wire
	identifier wires.Identifier
}

func (i AndInstruction) Apply() wires.Wire {
	return wires.NewWire(i.identifier, i.left.Signal & i.right.Signal)
}

func NewAndInstruction(left wires.Wire, right wires.Wire, identifier wires.Identifier) Instruction {
	return AndInstruction{
		left: left,
		right: right,
		identifier: identifier,
	}
}

type OrInstruction struct {
	left       wires.Wire
	right      wires.Wire
	identifier wires.Identifier
}

func (i OrInstruction) Apply() wires.Wire {
	return wires.NewWire(i.identifier, i.left.Signal | i.right.Signal)
}

func NewOrInstruction(left wires.Wire, right wires.Wire, identifier wires.Identifier) Instruction {
	return OrInstruction{
		left: left,
		right: right,
		identifier: identifier,
	}
}

type LShiftInstruction struct {
	wire       wires.Wire
	identifier wires.Identifier
	shift      wires.Shift
}

func (i LShiftInstruction) Apply() wires.Wire {
	return wires.NewWire(i.identifier, i.wire.Signal << i.shift)
}

func NewLShiftInstruction(wire wires.Wire, shift wires.Shift, identifier wires.Identifier) Instruction {
	return LShiftInstruction{
		wire: wire,
		shift: shift,
		identifier: identifier,
	}
}

type RShiftInstruction struct {
	wire       wires.Wire
	identifier wires.Identifier
	shift      wires.Shift
}

func (i RShiftInstruction) Apply() wires.Wire {
	return wires.NewWire(i.identifier, i.wire.Signal >> i.shift)
}

func NewRShiftInstruction(wire wires.Wire, shift wires.Shift, identifier wires.Identifier) Instruction {
	return RShiftInstruction{
		wire: wire,
		shift: shift,
		identifier: identifier,
	}
}

type NotInstruction struct {
	wire       wires.Wire
	identifier wires.Identifier
}

func (i NotInstruction) Apply() wires.Wire {
	return wires.NewWire(i.identifier, ^i.wire.Signal)
}

func NewNotInstruction(wire wires.Wire, identifier wires.Identifier) Instruction {
	return NotInstruction{
		wire: wire,
		identifier: identifier,
	}
}

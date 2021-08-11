package test

import (
	"day-7/internal/instructions"
	"day-7/internal/wires"
)

var xWire = wires.NewWire("x", 123)
var yWire = wires.NewWire("y", 456)

var applyInstructionTestCases = []struct {
	instruction  instructions.Instruction
	expectedWire wires.Wire
}{
	{
		instruction:  instructions.NewAndInstruction(xWire, yWire, "d"),
		expectedWire: wires.NewWire("d", 72),
	},
	{
		instruction: instructions.NewOrInstruction(xWire, yWire, "e"),
		expectedWire:    wires.NewWire("e", 507),
	},
	{
		instruction: instructions.NewLShiftInstruction(xWire, 2, "f"),
		expectedWire:    wires.NewWire("f", 492),
	},
	{
		instruction: instructions.NewRShiftInstruction(yWire, 2, "g"),
		expectedWire:    wires.NewWire("g", 114),
	},
	{
		instruction: instructions.NewNotInstruction(xWire, "h"),
		expectedWire:    wires.NewWire("h", 65412),
	},
	{
		instruction: instructions.NewNotInstruction(yWire, "i"),
		expectedWire:    wires.NewWire("i", 65079),
	},
}


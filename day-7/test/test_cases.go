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
		instruction:  instructions.NewOrInstruction(xWire, yWire, "e"),
		expectedWire: wires.NewWire("e", 507),
	},
	{
		instruction:  instructions.NewLShiftInstruction(xWire, 2, "f"),
		expectedWire: wires.NewWire("f", 492),
	},
	{
		instruction:  instructions.NewRShiftInstruction(yWire, 2, "g"),
		expectedWire: wires.NewWire("g", 114),
	},
	{
		instruction:  instructions.NewNotInstruction(xWire, "h"),
		expectedWire: wires.NewWire("h", 65412),
	},
	{
		instruction:  instructions.NewNotInstruction(yWire, "i"),
		expectedWire: wires.NewWire("i", 65079),
	},
}

var parseInstructionTestCases = []struct {
	rawInstruction      string
	expectedInstruction instructions.Instruction
}{
	{
		rawInstruction: "123 AND 456 -> a",
		expectedInstruction: instructions.NewAndInstruction(
			wires.Signal(123),
			wires.Signal(456),
			"a"),
	},
	{
		rawInstruction: "x AND 456 -> a",
		expectedInstruction: instructions.NewAndInstruction(
			wires.NewWire("x", 0),
			wires.Signal(456),
			"a"),
	},
	{
		rawInstruction: "123 AND y -> a",
		expectedInstruction: instructions.NewAndInstruction(
			wires.Signal(123),
			wires.NewWire("y", 0),
			"a"),
	},
	{
		rawInstruction: "x OR y -> b",
		expectedInstruction: instructions.NewOrInstruction(
			wires.NewWire("x", 0),
			wires.NewWire("y", 0),
			"b"),
	},
	{
		rawInstruction: "123 LSHIFT 2 -> c",
		expectedInstruction: instructions.NewLShiftInstruction(
			wires.Signal(123),
			2,
			"c"),
	},
	{
		rawInstruction: "x RSHIFT 3 -> d",
		expectedInstruction: instructions.NewRShiftInstruction(
			wires.NewWire("x", 0),
			3,
			"d"),
	},
	{
		rawInstruction: "NOT x -> e",
		expectedInstruction: instructions.NewNotInstruction(
			wires.NewWire("x", 0),
			"e"),
	},
}

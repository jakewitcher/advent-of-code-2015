package test

import "day-7/internal/domain"

var xWire = domain.NewWire("x", 123)
var yWire = domain.NewWire("y", 456)

var andOrTestCases = []struct {
	leftInputWire   domain.Wire
	rightInputWire  domain.Wire
	inputIdentifier domain.Identifier
	gate            domain.Gate
	expectedWire    domain.Wire
}{
	{
		leftInputWire:   xWire,
		rightInputWire:  yWire,
		inputIdentifier: "d",
		gate:            domain.And,
		expectedWire:    domain.NewWire("d", 72),
	},
	{
		leftInputWire:   xWire,
		rightInputWire:  yWire,
		inputIdentifier: "e",
		gate:            domain.Or,
		expectedWire:    domain.NewWire("e", 507),
	},
}

var lShiftRShiftTestCases = []struct {
	inputWire       domain.Wire
	inputIdentifier domain.Identifier
	gate            domain.Gate
	shift           int
	expectedWire    domain.Wire
}{
	{
		inputWire:       xWire,
		inputIdentifier: "f",
		gate:            domain.LShift,
		shift:           2,
		expectedWire:    domain.NewWire("f", 492),
	},
	{
		inputWire:       yWire,
		inputIdentifier: "g",
		gate:            domain.RShift,
		shift:           2,
		expectedWire:    domain.NewWire("g", 114),
	},
}

var notTestCases = []struct {
	inputWire       domain.Wire
	inputIdentifier domain.Identifier
	gate            domain.Gate
	expectedWire    domain.Wire
}{
	{
		inputWire:       xWire,
		inputIdentifier: "h",
		gate:            domain.Not,
		expectedWire:    domain.NewWire("h", 65412),
	},
	{
		inputWire:       yWire,
		inputIdentifier: "i",
		gate:            domain.Not,
		expectedWire:    domain.NewWire("i", 65079),
	},
}

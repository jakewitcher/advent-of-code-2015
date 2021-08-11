package parse

import (
	"day-7/internal/instructions"
	"day-7/internal/wires"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	and    = "AND"
	or     = "OR"
	lShift = "LSHIFT"
	rShift = "RSHIFT"
	not    = "NOT"
)

var regex = regexp.MustCompile("(\\d+|[a-z])")

func All(rawInstructions []string) ([]instructions.Instruction, error) {
	parsedInstructions := make([]instructions.Instruction, len(rawInstructions))

	for i, rawInstruction := range rawInstructions {
		instruction, err := One(rawInstruction)
		if err != nil {
			return nil, err
		}

		parsedInstructions[i] = instruction
	}

	return parsedInstructions, nil
}

func One(rawInstruction string) (instructions.Instruction, error) {
	var instruction instructions.Instruction

	switch {
	case strings.Contains(rawInstruction, and):
		instruction = parseAndOrInstruction(rawInstruction)

	case strings.Contains(rawInstruction, or):
		instruction = parseOrInstruction(rawInstruction)

	case strings.Contains(rawInstruction, lShift):
		instruction = parseLShiftInstruction(rawInstruction)

	case strings.Contains(rawInstruction, rShift):
		instruction = parseRShiftInstruction(rawInstruction)

	case strings.Contains(rawInstruction, not):
		instruction = parseNotInstruction(rawInstruction)
	}

	if instruction == nil {
		return instruction, fmt.Errorf("invalid instruction: %s", rawInstruction)
	}

	return instruction, nil
}

func parseAndOrInstruction(rawInstruction string) instructions.Instruction {
	match := regex.FindAll([]byte(rawInstruction), 3)

	left := parseWire(string(match[0]))
	right := parseWire(string(match[1]))
	identifier := wires.Identifier(match[2])

	return instructions.NewAndInstruction(left, right, identifier)
}

func parseOrInstruction(rawInstruction string) instructions.Instruction {
	match := regex.FindAll([]byte(rawInstruction), 3)

	left := parseWire(string(match[0]))
	right := parseWire(string(match[1]))
	identifier := wires.Identifier(match[2])

	return instructions.NewOrInstruction(left, right, identifier)
}

func parseLShiftInstruction(rawInstruction string) instructions.Instruction {
	match := regex.FindAll([]byte(rawInstruction), 3)

	wire := parseWire(string(match[0]))
	shift := parseShift(string(match[1]))
	identifier := wires.Identifier(match[2])

	return instructions.NewLShiftInstruction(wire, shift, identifier)
}

func parseRShiftInstruction(rawInstruction string) instructions.Instruction {
	match := regex.FindAll([]byte(rawInstruction), 3)

	wire := parseWire(string(match[0]))
	shift := parseShift(string(match[1]))
	identifier := wires.Identifier(match[2])

	return instructions.NewRShiftInstruction(wire, shift, identifier)
}

func parseNotInstruction(rawInstruction string) instructions.Instruction {
	match := regex.FindAll([]byte(rawInstruction), 2)

	wire := parseWire(string(match[0]))
	identifier := wires.Identifier(match[1])

	return instructions.NewNotInstruction(wire, identifier)
}

func parseWire(rawWire string) wires.Wire {
	var identifier wires.Identifier
	var signal wires.Signal

	maybeSignal, err := strconv.Atoi(rawWire)
	if err != nil {
		identifier = wires.Identifier(rawWire)
	} else {
		signal = wires.Signal(maybeSignal)
	}

	return wires.NewWire(identifier, signal)
}

func parseShift(rawShift string) wires.Shift {
	shift, _ := strconv.Atoi(rawShift)
	return wires.Shift(shift)
}

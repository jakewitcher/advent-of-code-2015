package elevator

import (
	"fmt"
)

func Run(route string) int {
	var currentFloor int

	for _, r := range route {
		switch r {
		case '(':
			currentFloor++
		case ')':
			currentFloor--
		}
	}

	return currentFloor
}

func FindPosition(route string, floor int) (int, error) {
	var currentFloor int

	for i, r := range route {
		switch r {
		case '(':
			currentFloor++
		case ')':
			currentFloor--
		}

		if currentFloor == floor {
			return i + 1, nil
		}
	}

	return 0, fmt.Errorf("expected floor %d not found", floor)
}

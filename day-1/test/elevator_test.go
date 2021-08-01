package elevator

import (
	"day-1/internal/elevator"
	"testing"
)

func TestElevatorRun(t *testing.T) {
	for _, test := range elevatorRunTestCases {
		if floor := elevator.Run(test.route); floor != test.floor {
			t.Fatalf("floor floor: %d, floor floor: %d", test.floor, floor)
		}
	}
}

func TestElevatorFindPosition(t *testing.T) {
	for _, test := range elevatorPositionTestCases {
		if position, _ := elevator.FindPosition(test.route, test.floor); position != test.position {
			t.Fatalf("floor position: %d, position position: %d", test.position, position)
		}
	}
}

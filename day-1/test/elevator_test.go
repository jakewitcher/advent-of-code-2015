package test

import (
	"day-1/internal/elevator"
	"day-1/internal/input"
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

func BenchmarkElevatorRun(b *testing.B) {
	route, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = elevator.Run(route)
	}
}

func BenchmarkElevatorFindPosition(b *testing.B) {
	route, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = elevator.FindPosition(route, -1)
	}
}
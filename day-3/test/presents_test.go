package test

import (
	"day-3/internal/input"
	"day-3/internal/presents"
	"testing"
)

func TestPresentsDeliveredBySanta(t *testing.T) {
	for _, test := range santaDeliveryTestCases {
		if houses := presents.DeliveredBySanta(test.route); houses != test.houses {
			t.Fatalf("expected number of house: %d, actual number of house: %d", test.houses, houses)
		}
	}
}

func TestPresentsDeliveredBySantaAndRobotSanta(t *testing.T) {
	for _, test := range robotSantaDeliveryTestCases {
		if houses := presents.DeliveredBySantaAndRobotSanta(test.route); houses != test.houses {
			t.Fatalf("expected number of house: %d, actual number of house: %d", test.houses, houses)
		}
	}
}

func BenchmarkPresentsDeliveredBySanta(b *testing.B) {
	route, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = presents.DeliveredBySanta(route)
	}
}

func BenchmarkPresentsDeliveredBySantaAndRobotSanta(b *testing.B) {
	route, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = presents.DeliveredBySantaAndRobotSanta(route)
	}
}

package test

import (
	"day-4/internal/input"
	"day-4/internal/mining"
	"log"
	"testing"
)

func TestFindHashWithPrefix(t *testing.T) {
	for _, test := range testCases {
		if number := mining.FindHashWithPrefix("00000", test.secret); number != test.number {
			t.Fatalf("expected number: %d, actual number: %d", test.number, number)
		}
	}
}

func BenchmarkFindHashWithPrefix(b *testing.B) {
	secretKey, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = mining.FindHashWithPrefix("000000", secretKey)
	}
}

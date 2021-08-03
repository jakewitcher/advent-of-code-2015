package main

import (
	"day-4/internal/input"
	"day-4/internal/mining"
	"log"
)

func main() {
	secretKey, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	PartOne(secretKey)
	PartTwo(secretKey)
}

func PartOne(secretKey string) {
	prefix := "00000"

	number := mining.FindHashWithPrefix(prefix, secretKey)
	log.Printf("number: %d", number)
}

func PartTwo(secretKey string) {
	prefix := "000000"

	number := mining.FindHashWithPrefix(prefix, secretKey)
	log.Printf("number: %d", number)
}

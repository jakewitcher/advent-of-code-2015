package main

import (
	"day-3/internal/input"
	"day-3/internal/presents"
	"log"
)

func main() {
	route, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	PartOne(route)
	PartTwo(route)
}

func PartOne(route string) {
	houses := presents.DeliveredBySanta(route)
	log.Printf("houses visited: %d", houses)
}

func PartTwo(route string) {
	houses := presents.DeliveredBySantaAndRobotSanta(route)
	log.Printf("houses visited: %d", houses)
}

package main

import (
	"day-1/internal/elevator"
	"day-1/internal/input"
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
	floor := elevator.Run(route)
	log.Printf("floor: %d", floor)
}

func PartTwo(route string) {
	position, err := elevator.FindPosition(route, -1)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("position: %d", position)
}
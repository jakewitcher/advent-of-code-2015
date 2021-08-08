package main

import (
	"day-6/internal/domain"
	"day-6/internal/input"
	"day-6/internal/lights"
	"log"
)

func main() {
	directions, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	PartOne(directions)
	PartTwo(directions)
}

func PartOne(directions []string) {
	lightsOn := lights.WithDirections(directions, domain.NewBinaryLight)
	log.Printf("lights on: %d", lightsOn)
}

func PartTwo(directions []string) {
	lightsOn := lights.WithDirections(directions, domain.NewBrightnessLight)
	log.Printf("lights on: %d", lightsOn)
}
package main

import (
	"day-6/internal/grid"
	"day-6/internal/input"
	"day-6/internal/light"
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
	lightsOn := grid.LightAccordingToDirections(directions, light.NewBinaryLight)
	log.Printf("grid on: %d", lightsOn)
}

func PartTwo(directions []string) {
	lightsOn := grid.LightAccordingToDirections(directions, light.NewBrightnessLight)
	log.Printf("grid on: %d", lightsOn)
}
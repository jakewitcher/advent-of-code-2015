package main

import (
	"day-6/internal/input"
	"day-6/internal/lights"
	"log"
)

func main() {
	directions, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lightsOn := lights.WithDirections(directions)
	log.Printf("lights on: %d", lightsOn)
}

package main

import (
	"day-2/internal/input"
	"day-2/internal/presents"
	"log"
)

func main() {
	dimensions, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	PartOne(dimensions)
	PartTwo(dimensions)
}

func PartOne(dimensions []string) {
	sqFt, err := presents.CalculateWrappingPaperRequired(dimensions)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("total sq ft of wrapping presents: %d", sqFt)
}

func PartTwo(dimensions []string) {
	ft, err := presents.CalculateRibbonRequired(dimensions)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("total sq ft of wrapping presents: %d", ft)
}

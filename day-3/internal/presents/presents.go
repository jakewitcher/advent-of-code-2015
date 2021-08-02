package presents

import "day-3/internal/domain"

type Houses = map[domain.Point]struct{}

func DeliveredBySanta(route string) int {
	houses := make(Houses)
	current := domain.NewPoint(0, 0)

	houses[current] = struct{}{}
	for _, r := range route {
		current = getNextHouse(r, current)
		houses[current] = struct{}{}
	}

	return len(houses)
}

func DeliveredBySantaAndRobotSanta(route string) int {
	houses := make(Houses)
	santaCurrent := domain.NewPoint(0, 0)
	robotCurrent := domain.NewPoint(0, 0)

	houses[santaCurrent] = struct{}{}
	for i, r := range route {
		if i%2 == 0 {
			robotCurrent = getNextHouse(r, robotCurrent)
			houses[robotCurrent] = struct{}{}
		} else {
			santaCurrent = getNextHouse(r, santaCurrent)
			houses[santaCurrent] = struct{}{}
		}
	}

	return len(houses)
}

func getNextHouse(r rune, current domain.Point) domain.Point {
	var next domain.Point

	switch r {
	case '^':
		next = current.Up()
	case '>':
		next = current.Right()
	case 'v':
		next = current.Down()
	case '<':
		next = current.Left()
	}

	return next
}

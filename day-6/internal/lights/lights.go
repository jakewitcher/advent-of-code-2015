package lights

import (
	"day-6/internal/direction"
	"day-6/internal/domain"
)

func WithDirections(rawDirections []string, lightFactory domain.LightFactory) int {
	directions := make([]domain.Direction, len(rawDirections))
	for i, d := range rawDirections {
		directions[i] = direction.Parse(d)
	}

	grid := domain.NewGrid(1000, 1000, lightFactory)

	for _, d := range directions {
		grid.ApplyDirection(d)
	}

	return grid.CountLightsOn()
}

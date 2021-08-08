package grid

import (
	d "day-6/internal/direction"
	l "day-6/internal/light"
)

type Grid [][]l.Light

func (g Grid) ApplyDirection(direction d.Direction) {
	for i := direction.Start.X; i <= direction.End.X; i++ {
		for j := direction.Start.Y; j <= direction.End.Y; j++ {
			switch direction.Action {
			case d.TurnOn:
				light := g[i][j]
				g[i][j] = light.On()
			case d.TurnOff:
				light := g[i][j]
				g[i][j] = light.Off()
			case d.Toggle:
				light := g[i][j]
				g[i][j] = light.Toggle()
			}
		}
	}
}

func (g Grid) CountLightsOn() int {
	var lightsOn int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			light := g[i][j]
			lightsOn += light.Status()
		}
	}

	return lightsOn
}

func NewGrid(width, height int, lightFactory l.Factory) Grid {
	grid := make(Grid, height)

	for i := 0; i < height; i++ {
		row := make([]l.Light, width)

		for j := 0; j < width; j++ {
			row[j] = lightFactory(i, j)
		}

		grid[i] = row
	}

	return grid
}

func LightAccordingToDirections(rawDirections []string, lightFactory l.Factory) int {
	directions := make([]d.Direction, len(rawDirections))
	for i, rawDirection := range rawDirections {
		directions[i] = d.Parse(rawDirection)
	}

	grid := NewGrid(1000, 1000, lightFactory)

	for _, d := range directions {
		grid.ApplyDirection(d)
	}

	return grid.CountLightsOn()
}

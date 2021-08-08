package test

import (
	d "day-6/internal/direction"
	"day-6/internal/light"
)

var parseDirectionTestCases = []struct {
	rawDirection string
	direction d.Direction
}{
	{
		rawDirection: "turn on 489,959 through 759,964",
		direction: d.Direction{
			Start: light.Position{
				X: 489,
				Y: 959,
			},
			End: light.Position{
				X: 759,
				Y: 964,
			},
			Action: d.TurnOn,
		},
	},
	{
		rawDirection: "turn off 820,516 through 871,914",
		direction: d.Direction{
			Start: light.Position{
				X: 820,
				Y: 516,
			},
			End: light.Position{
				X: 871,
				Y: 914,
			},
			Action: d.TurnOff,
		},
	},
	{
		rawDirection: "toggle 756,965 through 812,992",
		direction: d.Direction{
			Start: light.Position{
				X: 756,
				Y: 965,
			},
			End: light.Position{
				X: 812,
				Y: 992,
			},
			Action: d.Toggle,
		},
	},
}

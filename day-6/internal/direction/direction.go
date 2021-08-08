package direction

import (
	"day-6/internal/light"
	"regexp"
	"strconv"
	"strings"
)

const (
	TurnOn  string = "turn on"
	TurnOff string = "turn off"
	Toggle  string = "toggle"
)

type Direction struct {
	Start  light.Position
	End    light.Position
	Action string
}

var regex = regexp.MustCompile("\\d+")

func Parse(direction string) Direction {
	match := regex.FindAllString(direction, 4)
	startX, _ := strconv.Atoi(match[0])
	startY, _ := strconv.Atoi(match[1])

	endX, _ := strconv.Atoi(match[2])
	endY, _ := strconv.Atoi(match[3])

	action := parseAction(direction)

	return Direction{
		Start: light.Position{
			X: startX,
			Y: startY,
		},
		End: light.Position{
			X: endX,
			Y: endY,
		},
		Action: action,
	}
}

func parseAction(direction string) string {
	var action string

	if strings.HasPrefix(direction, TurnOn) {
		action = TurnOn
	} else if strings.HasPrefix(direction, TurnOff) {
		action = TurnOff
	} else {
		action = Toggle
	}

	return action
}

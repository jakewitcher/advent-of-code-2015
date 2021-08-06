package direction

import (
	"day-6/internal/domain"
	"regexp"
	"strconv"
	"strings"
)

const (
	turnOn  string = "turn on"
	turnOff string = "turn off"
	toggle  string = "toggle"
)

var regex = regexp.MustCompile("\\d+")

func Parse(direction string) domain.Direction {
	match := regex.FindAllString(direction, 4)
	startX, _ := strconv.Atoi(match[0])
	startY, _ := strconv.Atoi(match[1])

	endX, _ := strconv.Atoi(match[2])
	endY, _ := strconv.Atoi(match[3])

	action := parseAction(direction)

	return domain.Direction{
		Start: domain.Position{
			X: startX,
			Y: startY,
		},
		End: domain.Position{
			X: endX,
			Y: endY,
		},
		Action: action,
	}
}

func parseAction(direction string) string {
	var action string

	if strings.HasPrefix(direction, turnOn) {
		action = turnOn
	} else if strings.HasPrefix(direction, turnOff) {
		action = turnOff
	} else {
		action = toggle
	}

	return action
}

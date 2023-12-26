package client

import (
	"fmt"
	"strconv"
	"time"
)

type counter struct {
	count int
}

var Counter = counter{}

func (c *counter) TickCounter() {
	for {
		time.Sleep(1 * time.Second)
		var viewCounter = c.count
		c.count = 0
		divisor := 1024
		var unit = getNextUnit("")
		for viewCounter > divisor && getNextUnit(unit) != "B" {
			viewCounter /= divisor
			unit = getNextUnit(unit)
		}

		fmt.Println(strconv.Itoa(viewCounter) + unit)
	}
}

func (c *counter) Add(value int) {
	c.count += value
}

func getNextUnit(currentUnit string) string {
	switch currentUnit {
	case "B":
		return "KB"
	case "KB":
		return "MB"
	default:
		return "B"
	}
}

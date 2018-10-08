package main

import (
	"fmt"
)

func main() {
	println(north)
	commands := "GLGLGLGR"

	circleExists(commands)
}

const (
	north = iota
	east
	south
	west
)

type pos struct {
	x         int
	y         int
	direction int
}

func (p pos) findNewPosition(commands string) pos {
	for _, c := range commands {
		command := string(c)
		p = p.applyCommand(command)
	}

	return p
}

func (p pos) applyCommand(command string) pos {
	switch command {
	case "G":
		p.x, p.y = p.step()
	case "L":
		p.direction = p.turnLeft()
	default:
		p.direction = p.turnRight()
	}

	return p
}

func (p pos) step() (int, int) {
	switch p.direction {
	case north:
		p.y++
	case east:
		p.x++
	case south:
		p.y--
	default:
		p.x--
	}

	return p.x, p.y
}

func (p pos) turnLeft() int {
	switch p.direction {
	case north:
		return west
	case east:
		return north
	case south:
		return east
	default:
		return south
	}
}

func (p pos) turnRight() int {
	switch p.direction {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	default:
		return north
	}
}

func circleExists(commands string) bool {
	startPos := pos{}
	newPos := startPos

	for i := 0; i < 4; i++ {
		newPos = newPos.findNewPosition(commands)
	}

	fmt.Printf("start: %v\n", startPos)
	fmt.Printf("end: %v\n", newPos)

	return true
}

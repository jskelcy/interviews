package main

import (
	"fmt"
)

type coords struct {
	x int
	y int
}

func main() {
	islandMap := [][]int{
		{1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{1, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1},
	}

	fmt.Println(scanMap(islandMap))
}

func scanMap(islandMap [][]int) int {
	var islandCounter int
	visited := map[coords]bool{}

	for y, row := range islandMap {
		for x, cell := range row {
			currCell := coords{y: y, x: x}
			if !visited[currCell] && cell != 0 {
				markIsland(islandMap, visited, currCell)
				islandCounter++
			}
		}
	}

	return islandCounter
}

func markIsland(islandMap [][]int, visited map[coords]bool, currPos coords) {
	visited[currPos] = true

	// Check for neighbors
	neighborIndex := []int{-1, 0, 1}
	nieghbors := []coords{}
	for _, y := range neighborIndex {
		for _, x := range neighborIndex {
			neighborCord := coords{y: currPos.y + y, x: currPos.x + x}
			// value is on map
			if (neighborCord.y > -1 && neighborCord.x > -1) &&
				(neighborCord.y < len(islandMap) && neighborCord.x < len(islandMap[neighborCord.y])) {
				// value had not been visited or water
				if !visited[neighborCord] && islandMap[neighborCord.y][neighborCord.x] == 1 {
					nieghbors = append(nieghbors, neighborCord)
				}
			}
		}
	}

	// recurse on neighbors
	for _, nieghbor := range nieghbors {
		markIsland(islandMap, visited, nieghbor)
	}
}

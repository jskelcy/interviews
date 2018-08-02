package main

import (
	"fmt"
)

func main() {
	image := [][]int{
		[]int{1, 1, 1, 1, 1, 1, 1, 0},
		[]int{1, 0, 1, 1, 1, 1, 1, 0},
		[]int{1, 1, 1, 0, 0, 0, 1, 1},
		[]int{1, 1, 1, 0, 0, 0, 1, 1},
		[]int{0, 0, 1, 1, 1, 1, 1, 0},
	}

	foundRecs := findRecs(image)
	for _, rec := range foundRecs.r {
		fmt.Printf("topx: %v, topy: %v, botx: %v, boty: %v,\n", rec.topx, rec.topy, rec.botx, rec.boty)
	}
}

type rec struct {
	topx int
	topy int
	botx int
	boty int
}

func (r *rec) pointInRec(x, y int) bool {
	if (x >= r.topx && x <= r.botx) && (y >= r.topy && y <= r.boty) {
		return true
	}
	return false
}

type recs struct {
	r []rec
}

func (r *recs) pointInRecs(x, y int) bool {
	for _, rec := range r.r {
		if rec.pointInRec(x, y) {
			return true
		}
	}

	return false
}

func findRecs(image [][]int) recs {
	foundRecs := recs{r: []rec{}}

	for yCord, row := range image {
		for xCord, value := range row {
			if value == 0 {
				if !foundRecs.pointInRecs(xCord, yCord) {
					width, height := size(image, xCord, yCord)
					foundRecs.r = append(foundRecs.r, rec{
						topx: xCord,
						topy: yCord,
						botx: xCord + width - 1,
						boty: yCord + height - 1,
					})
				}
			}
		}
	}

	return foundRecs
}

func size(image [][]int, startX, startY int) (int, int) {
	var width, height int
	for i := startY; i < len(image); i++ {
		if image[i][startX] == 0 {
			height++
		} else {
			break
		}
	}
	for i := startX; i < len(image[startY]); i++ {
		if image[startY][i] == 0 {
			width++
		} else {
			break
		}
	}

	return width, height
}

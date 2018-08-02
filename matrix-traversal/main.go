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

type point struct {
	x int
	y int
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

	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] == 0 {
				if !foundRecs.pointInRecs(j, i) {
					width, height := size(image, i, j)
					foundRecs.r = append(foundRecs.r, rec{
						topx: j,
						topy: i,
						botx: j + width - 1,
						boty: i + height - 1,
					})
				}
			}
		}
	}

	return foundRecs
}

func size(image [][]int, startY, startX int) (int, int) {
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

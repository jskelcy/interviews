package main

import (
	"fmt"
)

func main() {
	board := [][]string{
		[]string{"G", "I", "Z"},
		[]string{"U", "E", "K"},
		[]string{"Q", "S", "E"},
	}

	bog := newBoggle(board)
	wordList := map[string]struct{}{}
	depthFirst(&bog, wordList, "", bog.board[0][0], []*cell{bog.board[0][0]})
	_, ok := wordList["GEEK"]
	if !ok {
		fmt.Println("GEEK not found")
	} else {
		fmt.Println("GEEK found")
	}

	// dict := []string{"GEEKS", "FOR", "QUIZ", "GO"}
	// found := []string{}
	// for _, word := range words {
	// 	for _, dictWord := range dict {
	// 		if dictWord == word {
	// 			found = append(found, word)
	// 		}
	// 	}
	// }

	// fmt.Printf("found: %s\n", strings.Join(found, ", "))
}

func depthFirst(bog *boggle, wordList map[string]struct{}, word string, c *cell, used []*cell) {
	wordList[word+c.value] = struct{}{}
	neighbors := findNeighbors(c.y, c.x, used, bog)
	for _, neighbor := range neighbors {
		used = append(used, neighbor)
		depthFirst(bog, wordList, word+c.value, neighbor, used)
		used = used[:len(used)-1]
	}

	return
}

func findNeighbors(y, x int, used []*cell, bog *boggle) []*cell {
	neighbors := []*cell{}
	for _, looky := range []int{-1, 0, 1} {
		for _, lookx := range []int{-1, 0, 1} {
			if (y+looky > -1 && x+lookx > -1) && (y+looky < len(bog.board) && x+lookx < len(bog.board)) {
				neighbor := bog.board[y+looky][x+lookx]
				if !in(used, neighbor) {
					neighbors = append(neighbors, neighbor)
				}
			}
		}
	}

	return neighbors
}

type cell struct {
	visited bool
	value   string
	x       int
	y       int
}
type boggle struct {
	board [][]*cell
}

func (b *boggle) clear() {
	for _, row := range b.board {
		for _, cell := range row {
			cell.visited = false
		}
	}
}

func newBoggle(board [][]string) boggle {
	boggle := boggle{
		board: make([][]*cell, 0, len(board)),
	}
	for y, row := range board {
		boggleRow := make([]*cell, 0, len(row))
		for x, value := range row {
			boggleRow = append(boggleRow, &cell{
				value: value,
				x:     x,
				y:     y,
			})
		}
		boggle.board = append(boggle.board, boggleRow)
	}

	return boggle
}

func in(used []*cell, target *cell) bool {
	for _, cell := range used {
		if target == cell {
			return true
		}
	}

	return false
}

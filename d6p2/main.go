package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// filepath := "input.txt"
	filepath := "test_input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	Map := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		Map = append(Map, []rune(line))
	}

	guard := Guard{}
	guard.orientation = 0
	guard.distinctPosition = make(map[MapPoint]bool)
	startingPoint := MapPoint{0, 0, 0}

	// Find guard starting pos
out:
	for i, row := range Map {
		for j, r := range row {
			if r == rune('^') {
				guard.curPosX = i
				guard.curPosY = j
				startingPoint = MapPoint{i, j, 0}
				guard.distinctPosition[startingPoint] = true
				break out
			}
		}
	}

	// Start moving guard by placing obstacles one at a time
	xLimit := len(Map)
	yLimit := len(Map[0])
	for x := 0; x < len(Map); x++ {
		for y := 0; y < len(Map[x]); y++ {
			// At this level we have the coordinates for our newly placed obstacle
			for {
				// Guard has exited the grid
				if !isPosInLimits(guard.curPosX, guard.curPosY, xLimit, yLimit) {
					break
				}
				i, j := guard.NextMove()
				if isPosInLimits(i, j, xLimit, yLimit) && Map[i][j] == rune('#') {
					fmt.Println("obstacle", i, j)
					guard.Rotate()
				}
			}
		}
	}

	for {
		// Guard has exited our map so we return the result
		if !isPosInLimits(guard.curPosX, guard.curPosY, xLimit, yLimit) {
			// We deduct one cause we don't want the last position which is out of the grid
			continue
		}
		i, j := guard.NextMove()
		if isPosInLimits(i, j, xLimit, yLimit) && Map[i][j] == rune('#') {
			fmt.Println("obstacle", i, j)
			guard.Rotate()
		}
		guard.Move()
	}
}

func isPosInLimits(xPos int, yPos int, xLimit int, yLimit int) bool {
	if xPos < 0 || xPos >= xLimit || yPos < 0 || yPos >= yLimit {
		return false
	}
	return true
}

type MapPoint struct {
	x           int
	y           int
	orientation int
}

type Guard struct {
	curPosX   int
	curPosY   int
	totalStep int

	// 0 - up, 1 - right, 2 - down, 3 - left
	orientation      int
	xLimit           int
	yLimit           int
	distinctPosition map[MapPoint]bool
}

func (g *Guard) GetCopy() Guard {
	newGuard := g
	newGuard.distinctPosition = make(map[MapPoint]bool)
	return *newGuard
}

// Returns false if move has been done before
func (g *Guard) Move() bool {
	switch g.orientation {
	case 0:
		g.curPosX--
	case 1:
		g.curPosY++
	case 2:
		g.curPosX++
	case 3:
		g.curPosY--
	}
	g.totalStep++
	if g.distinctPosition[MapPoint{g.curPosX, g.curPosY, g.orientation}] == true {
		fmt.Println("detected loop")
		return false
	}
	g.distinctPosition[MapPoint{g.curPosX, g.curPosY, g.orientation}] = true
	return true
}

func (g *Guard) NextMove() (int, int) {
	switch g.orientation {
	case 0:
		return g.curPosX - 1, g.curPosY
	case 1:
		return g.curPosX, g.curPosY + 1
	case 2:
		return g.curPosX + 1, g.curPosY
	case 3:
		return g.curPosX, g.curPosY - 1
	}
	panic("Not a valid orientation")
}

func (g *Guard) Rotate() {
	g.orientation = (g.orientation + 1) % 4
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := "input.txt"
	// filepath := "test_input.txt"
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

	startingPoint := MapPoint{0, 0, 0}

	// Find guard starting pos
out:
	for i, row := range Map {
		for j, r := range row {
			if r == rune('^') {
				startingPoint = MapPoint{i, j, 0}
				break out
			}
		}
	}

	// Start moving guard by placing obstacles one at a time
	xLimit := len(Map)
	yLimit := len(Map[0])
	totalObstacles := 0
	for x := 0; x < len(Map); x++ {
		for y := 0; y < len(Map[x]); y++ {
			// At this level we have the coordinates for our newly placed obstacle
			if x == startingPoint.x && y == startingPoint.y {
				continue
			}
			if Map[x][y] == rune('#') {
				continue
			}
			guard := NewGuard(startingPoint.x, startingPoint.y)
			distPositions := make(map[MapPoint]bool)
			for {
				// Guard has exited the grid
				if !isPosInLimits(guard.curPosX, guard.curPosY, xLimit, yLimit) {
					break
				}
				i, j := guard.NextMove()
				if _, ok := distPositions[guard.CurrentMapPoint()]; ok {
					totalObstacles++
					break
				}
				distPositions[guard.CurrentMapPoint()] = true
				if isPosInLimits(i, j, xLimit, yLimit) && (Map[i][j] == rune('#') || (i == x && j == y)) {
					guard.Rotate()
					continue
				}
				guard.Move()
			}
		}
	}
	fmt.Println(totalObstacles)
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
	curPosX int
	curPosY int

	// 0 - up, 1 - right, 2 - down, 3 - left
	orientation int
}

func NewGuard(startX int, startY int) Guard {
	g := Guard{}
	g.curPosX = startX
	g.curPosY = startY
	return g
}

func (g *Guard) CurrentMapPoint() MapPoint {
	return MapPoint{g.curPosX, g.curPosY, g.orientation}
}

// Returns false if move has been done before
func (g *Guard) Move() {
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

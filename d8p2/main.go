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

	mat := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mat = append(mat, []rune(line))
	}

	fmt.Println(calculateAntinodes(mat))
}

func calculateAntinodes(mat [][]rune) int {
	antennas := []Point{}
	// Find all antennas
	for i, line := range mat {
		for j, rn := range line {
			if rn == '.' {
				continue
			}
			antennas = append(antennas, Point{rn, i, j})
		}
	}

	// For each antenna go and look for the similar ones and create an antinode for them
	antinodes := make(map[Point]bool)
	for _, curAnt := range antennas {
		for _, ant := range antennas {
			if curAnt == ant {
				continue
			}
			if curAnt.freq == ant.freq {
				points := generateAllAntinode(curAnt, ant, len(mat))
				for _, p := range points {
					p.freq = rune('.')
					antinodes[p] = true
				}
				// fmt.Println(curAnt, ant, Point{curAnt.freq, x, y})
			}
		}
	}

	return len(antinodes)
}

type Point struct {
	freq rune
	x    int
	y    int
}

func generateAllAntinode(startPoint Point, secondPoint Point, boundary int) []Point {
	res := []Point{}
	// initial points are antinodes themselves
	res = append(res, startPoint, secondPoint)
	pointA := startPoint
	pointB := secondPoint
	for {
		x, y := getAntinodeCoordinates(pointA, pointB)
		newPoint := Point{startPoint.freq, x, y}
		// Boundary check
		if newPoint.x < 0 || newPoint.x >= boundary || newPoint.y < 0 || newPoint.y >= boundary {
			break
		}
		// Subtistute points to continue further generation
		pointA = pointB
		pointB = newPoint
		res = append(res, newPoint)
	}
	return res
}

func getAntinodeCoordinates(startPoint Point, endPoint Point) (int, int) {
	resX := 0
	resY := 0

	if startPoint.x < endPoint.x {
		resX = endPoint.x + endPoint.x - startPoint.x
	} else if startPoint.x > endPoint.x {
		resX = endPoint.x - (startPoint.x - endPoint.x)
	}
	if startPoint.y < endPoint.y {
		resY = endPoint.y + endPoint.y - startPoint.y
	} else if startPoint.y > endPoint.y {
		resY = endPoint.y - (startPoint.y - endPoint.y)
	}
	return resX, resY
}

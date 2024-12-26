package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	filepath := "input.txt"
	// filepath := "test_input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mat := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, c := range line {
			conv, _ := strconv.Atoi(string(c))
			row = append(row, conv)
		}
		mat = append(mat, row)
	}

	for i, row := range mat {
		for j, num := range row {
			if num == 0 {
				trails := findTrails(mat, i, j)
			}
		}
	}

}

func findTrails(mat [][]int, i int, j int) {
	// i, j is our 0 starting point

	findNextStep(mat, i, j)
}

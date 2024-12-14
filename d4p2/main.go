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

	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	res := getMasOccurences(matrix)
	fmt.Println(res)
}

// go through the matrix, find all the A and check if they are surrounded my the M/S characters
func getMasOccurences(mat [][]rune) int {
	res := 0
	for i := 1; i < len(mat)-1; i++ {
		for j := 1; j < len(mat[i])-1; j++ {
			if mat[i][j] != rune('A') {
				continue
			}
			if HasMasAntiDiagonal(mat, i, j) && HasMasDiagonal(mat, i, j) {
				res++
			}
		}
	}
	return res
}

func HasMasDiagonal(mat [][]rune, row int, col int) bool {
	if mat[row+1][col+1] == rune('M') && mat[row-1][col-1] == rune('S') {
		return true
	}
	if mat[row+1][col+1] == rune('S') && mat[row-1][col-1] == rune('M') {
		return true
	}
	return false
}

func HasMasAntiDiagonal(mat [][]rune, row int, col int) bool {
	if mat[row-1][col+1] == rune('M') && mat[row+1][col-1] == rune('S') {
		return true
	}
	if mat[row-1][col+1] == rune('S') && mat[row+1][col-1] == rune('M') {
		return true
	}
	return false
}

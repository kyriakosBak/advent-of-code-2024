package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

	res := getXmasOccurences(matrix)
	fmt.Println(res)
}

func getXmasOccurences(mat [][]rune) int {
	occurences := 0
	word := "XMAS"
	length := len(word)
	subSlice := make([]rune, length)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			// check horizontally left->right
			if j <= len(mat[i])-length && areEqual(mat[i][j:j+length], word) {
				occurences++
			}
			// Check horizontally right->left
			if j >= length-1 {
				copy(subSlice, mat[i][j-length+1:j+1])
				fmt.Println(string(subSlice))

				slices.Reverse(subSlice)
				if areEqual(subSlice, word) {
					occurences++
				}
			}
			// Check vertically top->bottom
			if i <= len(mat)-length {
				subSlice = getColSubslice(mat, j, i, i+length)
				if areEqual(subSlice, word) {
					occurences++
				}
			}
			// Check vertically bottom->top
			if i >= length-1 {
				subSlice = getColSubslice(mat, j, i-length+1, i+1)
				slices.Reverse(subSlice)
				if areEqual(subSlice, word) {
					occurences++
				}
			}
			// Check diagonally top-left->bottom-right
			if i <= len(mat)-length && j <= len(mat[j])-length {
				subSlice = getMainDiagonalSubslice(mat, i, i+(length-1), j, j+(length-1))
				if areEqual(subSlice, word) {
					occurences++
				}
			}
			// Check diagonally bottom-right->top-left
			if i >= length-1 && j >= length-1 {
				subSlice = getMainDiagonalSubslice(mat, i-(length-1), i, j-(length-1), j)
				slices.Reverse(subSlice)
				if areEqual(subSlice, word) {
					occurences++
				}
			}
			// Check diagonally bottom-left->top-right
			if i >= length-1 && j <= len(mat[j])-length {
				subSlice = getAntiDiagonalSubslice(mat, i, i-(length-1), j, j+(length-1))
				if areEqual(subSlice, word) {
					occurences++
				}
			}
			// Check diagonally top-right->bottom-left
			if i <= len(mat)-length && j >= length-1 {
				subSlice = getAntiDiagonalSubslice(mat, i+(length-1), i, j-(length-1), j)
				slices.Reverse(subSlice)
				if areEqual(subSlice, word) {
					occurences++
				}
			}

		}
	}

	return occurences
}

func areEqual(runeWord []rune, strWord string) bool {
	return string(runeWord) == strWord
}

func getMainDiagonalSubslice(mat [][]rune, startRow int, endRow int, startCol int, endCol int) []rune {
	res := []rune{}
	i := startRow
	j := startCol
	for i <= endRow && j <= endCol {
		res = append(res, mat[i][j])
		i++
		j++
	}
	return res
}

func getAntiDiagonalSubslice(mat [][]rune, startRow int, endRow int, startCol int, endCol int) []rune {
	res := []rune{}
	i := startRow
	j := startCol
	for i >= endRow && j <= endCol {
		res = append(res, mat[i][j])
		i--
		j++
	}
	return res
}

func getColSubslice(mat [][]rune, col int, start int, end int) []rune {
	res := []rune{}

	if start > end {
		panic("start was bigger than end")
	}

	for i := start; i < end; i++ {
		res = append(res, mat[i][col])
	}

	return res
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// filepath := "input.txt"
	filepath := "test_input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	arr := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			conv, _ := strconv.Atoi(string(c))
			arr = append(arr, conv)
		}
	}
	blocks := createBlocks(arr)
	movedBlocks := moveBlocks(blocks)
	checksum := calculateChecksum(movedBlocks)
	fmt.Println(checksum)

}

func calculateChecksum(arr []int) int {
	sum := 0
	for i, num := range arr {
		if num == -1 {
			break
		}
		sum += i * num
	}
	return sum
}

func moveBlocks(arr []int) []int {
	// res := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		for j, num2 := range arr[0:i] {
			if num2 == -1 {
				arr[j] = arr[i]
				arr[i] = -1
			}
		}
	}
	return arr
}

func createBlocks(arr []int) []int {
	res := []int{}
	id := 0
	for i, num := range arr {
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				res = append(res, id)
			}
			id++
			continue
		}
		for j := 0; j < num; j++ {
			res = append(res, -1)
		}
	}

	return res
}

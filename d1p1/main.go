package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read file and put into two different arrays
	filepath := "input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	array1 := []int{}
	array2 := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), " ")

		converted, _ := strconv.Atoi(splitted[0])
		array1 = append(array1, converted)

		converted, _ = strconv.Atoi(splitted[len(splitted)-1])
		array2 = append(array2, converted)
	}

	// Sort arrays
	sort.Slice(array1, func(i, j int) bool {
		return array1[i] < array1[j]
	})

	sort.Slice(array2, func(i, j int) bool {
		return array2[i] < array2[j]
	})

	// Loop through arrays with index and add the absolute difference to the result
	result := 0
	for i := range array1 {
		result += absDiff(array1[i], array2[i])
	}
	fmt.Println(result)
}

func absDiff(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

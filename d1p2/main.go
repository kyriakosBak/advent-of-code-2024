package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// Go thorugh second list and calculate the occurences
	occurences := make(map[int]int)
	for _, elem := range array2 {
		occurences[elem]++
	}

	// Go through the first list and calculate the result based on the occurences
	result := 0
	for _, elem := range array1 {
		occurence := occurences[elem]
		result += elem * occurence
	}
	fmt.Println(result)
}

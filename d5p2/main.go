package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	filepath := "input.txt"
	// filepath := "test_input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rules := make(map[int][]int)
	updates := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Add rule
		if strings.Contains(line, "|") {
			splitted := strings.Split(line, "|")
			num1, _ := strconv.Atoi(splitted[0])
			num2, _ := strconv.Atoi(splitted[1])
			if rules[num1] == nil {
				rules[num1] = []int{}
			}
			rules[num1] = append(rules[num1], num2)
		}
		// Add update
		if strings.Contains(line, ",") {
			updates = append(updates, convertToSlice(line))
		}
	}

	res := 0
	for _, update := range updates {
		if isUpdateValid(update, rules) {
			continue
		}
		fixedUpdate := fixUpdate(update, rules)
		// Grab middle element and sum it
		res += fixedUpdate[len(fixedUpdate)/2]
	}
	fmt.Println(res)
}

func fixUpdate(update []int, rules map[int][]int) []int {
	for idx, num := range update {
		rule := rules[num]
		for i := 0; i <= idx; i++ {
			if slices.Contains(rule, update[i]) {
				// move element
				newSlice := slices.Insert(update, i, num)
				newSlice = RemoveIndex(newSlice, idx+1)
				return fixUpdate(newSlice, rules)
			}
		}
	}

	return update
}

func RemoveIndex(arr []int, index int) []int {
	res := make([]int, 0, len(arr)-1)
	res = append(res, arr[:index]...)
	res = append(res, arr[index+1:]...)
	return res
}

// For each item in the list, check if the previous numbers are in his rules
// This breaks the rules and is considered if invalid
func isUpdateValid(update []int, rules map[int][]int) bool {
	for idx, num := range update {
		rule := rules[num]
		for i := 0; i <= idx; i++ {
			if slices.Contains(rule, update[i]) {
				return false
			}
		}
	}
	return true
}

func convertToSlice(str string) []int {
	res := []int{}
	splitted := strings.Split(str, ",")
	for _, s := range splitted {
		conv, _ := strconv.Atoi(s)
		res = append(res, conv)
	}
	return res
}

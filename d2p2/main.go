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
	// Read file and tranform into arrays of "levels"
	filepath := "input.txt"
	// filepath := "test_input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reports := []Report{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reports = append(reports, *NewReport(scanner.Text()))
	}

	safe := 0
	for _, report := range reports {
		fmt.Println("initial array", report.levels)
		if IsSafe(report.levels, false) || IsSafe(report.levels[1:], true) {
			safe++
			fmt.Println("========================", report)
		}
	}
	fmt.Println(safe)
}

type Report struct {
	levels        []int
	increasing    bool
	reportSkipped bool
}

func NewReport(str string) *Report {
	r := new(Report)
	for _, elem := range strings.Split(str, " ") {
		num, _ := strconv.Atoi(elem)
		r.levels = append(r.levels, num)
	}

	return r
}

// We don't catch the case where we have to remove the first element
func IsSafe(arr []int, haveRemoved bool) bool {
	fmt.Println("arr inside IsSafe", arr)
	isIncreasing := false
	if arr[0] < arr[1] {
		isIncreasing = true
	}

	for i := 1; i < len(arr); i++ {
		if isIncreasing && arr[i-1] > arr[i] {
			if haveRemoved {
				return false
			}
			subArr1 := RemoveIndex(arr, i-1)
			subArr2 := RemoveIndex(arr, i)
			fmt.Println("increasing", arr, subArr1, subArr2)
			return IsSafe(subArr1, true) || IsSafe(subArr2, true)
		}
		if !isIncreasing && arr[i-1] < arr[i] {
			if haveRemoved {
				return false
			}
			subArr1 := RemoveIndex(arr, i-1)
			subArr2 := RemoveIndex(arr, i)
			fmt.Println("decreasing", i, arr, subArr1, subArr2)
			return IsSafe(subArr1, true) || IsSafe(subArr2, true)
		}
		levelDiff := absDiff(arr[i], arr[i-1])
		if levelDiff < 1 || levelDiff > 3 {
			if haveRemoved {
				return false
			}
			subArr1 := RemoveIndex(arr, i-1)
			subArr2 := RemoveIndex(arr, i)
			fmt.Println("diff", i, arr, subArr1, subArr2)
			return IsSafe(subArr1, true) || IsSafe(subArr2, true)
		}
	}

	return true
}

func RemoveIndex(arr []int, index int) []int {
	fmt.Println("index:", index)
	res := make([]int, 0, len(arr)-1)
	res = append(res, arr[:index]...)
	res = append(res, arr[index+1:]...)
	return res
}

func absDiff(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

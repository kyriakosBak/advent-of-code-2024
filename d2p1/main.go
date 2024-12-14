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
		if report.IsSafe() {
			safe++
		}
	}
	fmt.Println(safe)

}

type Report struct {
	levels     []int
	increasing bool
}

func NewReport(str string) *Report {
	r := new(Report)
	for _, elem := range strings.Split(str, " ") {
		num, _ := strconv.Atoi(elem)
		r.levels = append(r.levels, num)
	}
	if r.levels[0] < r.levels[1] {
		r.increasing = true
	}
	return r
}

func (r Report) IsSafe() bool {
	for i := 1; i < len(r.levels); i++ {
		if r.increasing && r.levels[i] < r.levels[i-1] {
			return false
		}
		if !r.increasing && r.levels[i] > r.levels[i-1] {
			return false
		}
		levelDiff := absDiff(r.levels[i], r.levels[i-1])
		if levelDiff < 1 || levelDiff > 3 {
			return false
		}
	}

	return true
}

func absDiff(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

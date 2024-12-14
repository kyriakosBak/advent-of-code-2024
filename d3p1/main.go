package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	filepath := "input.txt"
	// filepath := "test_input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	results := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results = append(results, extractMul(scanner.Text())...)
	}

	sum := 0
	for _, mul := range results {
		sum += getProduct(mul)
	}
	fmt.Println(sum)
}

// some of the results may contains double comas which we need to invalidate when parsing
func extractMul(str string) []string {
	res := []string{}

	mulString := "mul("
	// Find occurences of mul(
	for i := 0; i < len(str)-3; i++ {
		if str[i:i+4] == mulString {
			subStr := "mul("
			for _, j := range str[i+4:] {
				if unicode.IsDigit(j) || j == rune(',') {
					subStr += string(j)
					continue
				}
				if j == rune(')') {
					subStr += string(j)
					res = append(res, subStr)
					break
				}
				break
			}
		}
	}

	return res
}

func getProduct(mul string) int {
	if strings.Count(mul, ",") > 1 {
		return 0
	}
	numsWithComma := mul[4 : len(mul)-1]
	splitted := strings.Split(numsWithComma, ",")
	num1, _ := strconv.Atoi(splitted[0])
	num2, _ := strconv.Atoi(splitted[1])
	return num1 * num2
}

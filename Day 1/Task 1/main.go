package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// file, err := os.Open("Test.txt")
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//for each line, find the first and last didgit
	LineValues := make([]int, 0)
	for i, line := range lines {
		first := ""
		last := ""
		for _, char := range line {
			if char >= '0' && char <= '9' {
				if first == "" {
					first = string(char)
				}
				last = string(char)
			}
		}
		combination := first + last
		intResult, err := strconv.Atoi(combination)
		if err != nil {
			fmt.Println("error on line ", i, " with text ", line)
			fmt.Println(err)
			return
		}
		LineValues = append(LineValues, intResult)
	}

	//print the sum of all the first and last digits
	sum := 0
	for _, value := range LineValues {
		sum += value
	}
	fmt.Println(sum)
}

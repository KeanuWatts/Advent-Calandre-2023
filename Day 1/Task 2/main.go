package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var debug = false

func main() {

	// file, err := os.Open("Test.txt")
	// file, err := os.Open("Input.txt")
	file, err := os.Open("onne.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	ogLines := make([]string, 0)
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		if lineNumber == 989 {
			debug = true
		} else {
			debug = false
		}
		tokens := tokenize(scanner.Text())
		newLine := ""
		for _, token := range tokens {
			newLine += token
		}
		lines = append(lines, newLine)
		ogLines = append(ogLines, scanner.Text())

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
	// print the sum of all the first and last digits
	sum := 0
	for _, value := range LineValues {
		// fmt.Println(strconv.Itoa(i) + ": " + strconv.Itoa(value) + " " + ogLines[i] + " " + lines[i])
		fmt.Println(value)
		sum += value
	}
	fmt.Println(sum)
}

func tokenize(input string) []string {
	var tokens []string
	var currentToken strings.Builder
	debugPrint("tokenizing: ", input)

	//we want to identify the following tokens:
	// "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"
	// "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"
	//anything else can be ignored
	ValidTextTokens := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}
	TranslatedTokens := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	ClearTokensTo := []string{"e", "o", "e", "", "e", "", "n", "t", "e", "o"}
	for _, char := range input {
		debugPrint("current char: ", string(char))
		if unicode.IsDigit(char) {
			debugPrint("found digit: ", string(char))
			tokens = append(tokens, string(char))
			currentToken.Reset()
			continue
		}

		currentToken.WriteRune(char)
		var foundMatch bool
		for i, validTextToken := range ValidTextTokens {
			//first check if the current token contains a valid text token
			if strings.Contains(currentToken.String(), validTextToken) {
				debugPrint("found valid text token: ", validTextToken)
				tokens = append(tokens, TranslatedTokens[i])
				currentToken.Reset()
				currentToken.WriteString(ClearTokensTo[i])
			}
		}

		for _, validTextToken := range ValidTextTokens {
			//if we have a partial match, we need to keep going
			debugPrint("checking for partial match for: ", currentToken.String())
			if strings.HasPrefix(validTextToken, currentToken.String()) {
				debugPrint("found partial match: ", currentToken.String())
				foundMatch = true
				break
			}
		}

		//if we don't have a match, we can reset
		if !foundMatch {
			debugPrint("no match found, resetting token ", currentToken.String())
			token := currentToken.String()
			currentToken.Reset()
			for len(token) > 0 {
				for _, validTextToken := range ValidTextTokens {
					if strings.HasPrefix(validTextToken, token) {
						debugPrint("found partial match: ", token)
						foundMatch = true
						currentToken.WriteString(token)
						break
					}
				}
				if !foundMatch {
					token = token[1:]
				} else {
					break
				}
			}
		}
	}
	return tokens
}

func debugPrint(a ...any) (n int, err error) {
	if debug {
		return fmt.Println(a...)
	}
	return
}

//first attempt 54105
//second attempt 54095

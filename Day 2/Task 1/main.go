package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cubes struct {
	Color string `json:"color"`
	Count int    `json:"count"`
}

type Game struct {
	GameNumber  int     `json:"GameNumber"`
	GameResults []Cubes `json:"GameResults"`
}

func main() {

	// file, err := os.Open("Test.txt")
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	games := make([]Game, 0)
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		games = append(games, Game{i, ParseLine(scanner.Text())})
		i++
	}
	// DisplayStruct(games)
	GameSum := 0
	for _, game := range games {
		if IsPossible(game, []Cubes{{"red", 12}, {"green", 13}, {"blue", 14}}) {
			GameSum += game.GameNumber
		}
	}
	fmt.Println(GameSum)
}

func ParseLine(input string) []Cubes {
	//for each line would look like this Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	cubes := make([]Cubes, 0)
	input = input[strings.Index(input, ":")+1:]
	GameSplit := strings.Split(input, ";")
	for _, line := range GameSplit {
		//slpit on the comma
		ColorSplit := strings.Split(line, ",")
		for _, subline := range ColorSplit {
			//split on the space
			ProperetySplit := strings.Split(subline, " ")
			// fmt.Println(ProperetySplit)
			count, err := strconv.Atoi(ProperetySplit[1]) //leading space, so start at 1
			if err != nil {
				fmt.Println(err)
				return nil
			}
			cubes = append(cubes, Cubes{ProperetySplit[2], count})
		}
	}
	return cubes
}

func IsPossible(game Game, maxCubes []Cubes) bool {
	//we want to deetermien if any of the cubes in the game were more than the max cubes of that color
	// fmt.Println("Testing Game ", game)
	for _, cube := range game.GameResults {
		for _, maxCube := range maxCubes {
			if cube.Color == maxCube.Color && cube.Count > maxCube.Count {

				fmt.Println("Game", game.GameNumber, "is not possible because there are", cube.Count, cube.Color, "cubes and only", maxCube.Count, "are allowed")
				return false
			}
		}
	}
	return true
}

func DisplayStruct(s interface{}) {
	b, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(b))
}

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
	DisplayStruct(games)
	maxSets := make([]Game, 0)
	for _, game := range games {
		maxSets = append(maxSets, GetMaxSet(game))
	}
	DisplayStruct(maxSets)
	totalsum := 0
	for _, maxSet := range maxSets {
		totalsum += GetPowerSet(maxSet)
	}
	fmt.Println(totalsum)
}

func GetPowerSet(game Game) int {
	//multiply the count of each color together
	res := 1
	for _, cube := range game.GameResults {
		res *= cube.Count
	}
	return res
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

func GetMaxSet(game Game) Game {
	res := Game{game.GameNumber, make([]Cubes, 0)}
	for _, cube := range game.GameResults {
		//first check if the color is already in the result set
		if !ContainsColor(res.GameResults, cube.Color) {
			res.GameResults = append(res.GameResults, cube)
		} else {
			//find the color and if we have more than the current count, update it
			for i, resCube := range res.GameResults {
				if resCube.Count < cube.Count {

					if resCube.Color == cube.Color {
						res.GameResults[i].Count = cube.Count
					}
				}
			}
		}
	}
	return res
}

func ContainsColor(cubes []Cubes, color string) bool {
	for _, cube := range cubes {
		if cube.Color == color {
			return true
		}
	}
	return false
}

func DisplayStruct(s interface{}) {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(b))
}

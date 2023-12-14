package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := "/Users/kevincammarata/Documents/Code/AoC23/Day2/puzzle2input.dat"
	file, err := os.Open(filepath)
	check(err)

	// var redLimit int = 12
	// var greenLimit int = 13
	// var blueLimit int = 14
	var gameTotal, powerTotal int

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		// Tokenize line
		gameLine := strings.Split(fileScanner.Text(), ":")
		gameNumber := parseGameNumber(gameLine[0])

		if checkMaxColorValuesForImpossibility(gameLine[1]) {
			gameTotal += gameNumber
			// fmt.Printf("Valid gameNumber %d \n", gameNumber)
		}
		powerTotal += powerFunction(gameLine[1])

		// if good, add Game Number (which is index + 1) to total
	}

	file.Close()

	fmt.Printf("result %d \n", gameTotal)
	fmt.Printf("Total Power %d \n", powerTotal)
	// fmt.Printf("result with text %v \n", totalWithText)

}

func parseGameNumber(gameHeader string) int {
	gameText := strings.Split(gameHeader, " ")
	gameNum, err := strconv.Atoi(gameText[1])
	check(err)
	return gameNum
}

// Given one Game, evaluate multiple draws and decide if the game was impossible
func checkMaxColorValuesForImpossibility(games string) bool {
	var maxRed, maxGreen, maxBlue int = 12, 13, 14
	splitter := func(r rune) bool {
		return r == ';' || r == ','
	}
	gameList := strings.FieldsFunc(games, splitter)

	for _, game := range gameList {
		gameDetails := strings.Split(game, " ")
		var i int
		if gameDetails[0] == "" {
			i++
		}
		blocksString := strings.TrimSpace(gameDetails[i])
		blocks, err := strconv.Atoi(blocksString)
		check(err)
		gameColor := gameDetails[(i + 1)]
		if gameColor == "red" && blocks > maxRed {
			return false
		}
		if gameColor == "blue" && blocks > maxBlue {
			return false
		}
		if gameColor == "green" && blocks > maxGreen {
			return false
		}

	}
	return true //If nothing was impossible
}

// Given one Game, evaluate multiple draws and decide if the game was impossible
func powerFunction(games string) int {
	var maxRed, maxGreen, maxBlue int
	splitter := func(r rune) bool {
		return r == ';' || r == ','
	}
	gameList := strings.FieldsFunc(games, splitter)

	for _, game := range gameList {
		gameDetails := strings.Split(game, " ")
		var i int
		if gameDetails[0] == "" {
			i++
		}
		blocksString := strings.TrimSpace(gameDetails[i])
		blocks, err := strconv.Atoi(blocksString)
		check(err)
		gameColor := gameDetails[(i + 1)]
		if gameColor == "red" && blocks > maxRed {
			maxRed = blocks
		}
		if gameColor == "blue" && blocks > maxBlue {
			maxBlue = blocks
		}
		if gameColor == "green" && blocks > maxGreen {
			maxGreen = blocks
		}

	}
	return (maxBlue * maxGreen * maxRed)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

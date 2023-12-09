package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	data := fetchTestData("/Users/kevincammarata/Documents/Code/AoC23/Day1/puzzle1input.dat")

// 	var result = processCalibrationData(data)
// 	fmt.Print("result", result)
// }

func processCalibrationData(rawData string) int {
	var sum, firstPos, secondPos int
	var firstDigit, secondDigit string
	firstPos = len(rawData)
	textNumbers := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i := 1; i < 10; i++ {
		testValue := textNumbers[i]
		iterAsString := strconv.Itoa(i)
		firstIndex := strings.Index(rawData, testValue)
		firstNumIndex := strings.Index(rawData, iterAsString)
		lastIndex := strings.LastIndex(rawData, testValue)
		lastNumIndex := strings.LastIndex(rawData, iterAsString)
		if firstIndex > -1 && firstIndex < firstPos {
			firstDigit = iterAsString
			firstPos = firstIndex
		}
		if firstNumIndex > -1 && firstNumIndex < firstPos {
			firstDigit = iterAsString
			firstPos = firstNumIndex
		}
		if lastIndex > secondPos {
			secondDigit = iterAsString
			secondPos = lastIndex
		}
		if lastNumIndex > secondPos {
			secondDigit = iterAsString
			secondPos = lastNumIndex
		}
	}

	newNum := firstDigit + secondDigit
	var err error
	sum, err = strconv.Atoi(newNum)
	if err != nil {
		fmt.Sprintf("Could not convert %s to an int", newNum)
		panic(err)
	}

	return sum
}

func fetchTestData(filepath string) []string {

	file, err := os.Open(filepath)
	check(err)
	fileLines := make([]string, 1001)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()

	// for _, line := range fileLines {
	//     fmt.Println(line)
	// }

	return fileLines
}

func fetchTestDataFromWeb(url string) string {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	fmt.Println("Response status:", resp.Status)
	// fmt.Println("Body: ", response)

	return string(response)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

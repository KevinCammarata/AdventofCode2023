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

func main() {
	filepath := "/Users/kevincammarata/Documents/Code/AoC23/Day1/puzzle1input.dat"
	file, err := os.Open(filepath)
	check(err)
	// coordinates := make([]int, 0)
	var total, totalWithText int

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		// fileLines = append(fileLines, fileScanner.Text())
		num := processCalibrationDataDigitsOnly(fileScanner.Text())
		numWithText := processCalibrationData(fileScanner.Text())
		// coordinates = append(coordinates, num)
		total += num
		totalWithText += numWithText
	}

	file.Close()

	fmt.Printf("result %d \n", total)
	fmt.Printf("result with text %v \n", totalWithText)
}

func processCalibrationDataDigitsOnly(rawData string) int {
	// fmt.Printf("Input to processCalibrationDataDigitsOnly %s", rawData)
	var sum, firstPos, secondPos int
	var firstDigit, secondDigit string
	firstPos = len(rawData)

	for i := 1; i < 10; i++ {
		iterAsString := strconv.Itoa(i)
		firstNumIndex := strings.Index(rawData, iterAsString)
		lastNumIndex := strings.LastIndex(rawData, iterAsString)

		if firstNumIndex > -1 && firstNumIndex < firstPos {
			firstDigit = iterAsString
			firstPos = firstNumIndex
		}

		if lastNumIndex >= secondPos {
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
		if lastNumIndex >= secondPos {
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

func fetchTestData(filepath string) int {

	file, err := os.Open(filepath)
	check(err)
	fileLines := make([]string, 1000)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()

	coordinates := make([]int, 1000)
	var total int

	for i := 1; i < 1000; i++ {
		num := processCalibrationDataDigitsOnly(fileLines[i])
		coordinates = append(coordinates, num)
		total = total + num
	}

	return total
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

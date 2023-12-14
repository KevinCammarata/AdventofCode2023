package main

import (
	"bufio"
	"os"
)

func main() {
	filepath := "/Users/kevincammarata/Documents/Code/AoC23/Day2/puzzle2input.dat"
	file, err := os.Open(filepath)
	check(err)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		// do some stuff
	}

	file.Close()

	// fmt.Printf("result %d \n", total)
	// fmt.Printf("result with text %v \n", totalWithText)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

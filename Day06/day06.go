package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReturnInputFromFile(path string) []string {
	var InputArray []string
	f, _ := os.Open(path)

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		InputArray = append(InputArray, scanner.Text())
	}
	return InputArray
}

func removeElementFromSlice(index int, slice [14]string) []string {
	return append(slice[:index], slice[index+1:]...)
}

func testIfSliceHasDuplicates(slice [14]string) bool {
	var sliceWithoutCurrentElement []string
	for i, element := range slice {
		sliceWithoutCurrentElement = removeElementFromSlice(i, slice)
		for _, elementFromSliceWithoutCurrentElement := range sliceWithoutCurrentElement {
			if element == elementFromSliceWithoutCurrentElement {
				return true
			}
		}
	}
	return false
}

func main() {
	var FileInput string = ReturnInputFromFile("day06_input.txt")[0]
	var charArray [14]string
	var charArrayIndex int = 0
	var resultIndex int = 0
	var sliceInited bool = false

	for _, letter := range FileInput {
		charArray[charArrayIndex] = string(letter)
		charArrayIndex += 1
		resultIndex += 1

		if charArrayIndex > 13 {
			charArrayIndex = 0
			sliceInited = true
		}

		//fmt.Printf("Current Slice: %v \n", charArray)
		if sliceInited && !testIfSliceHasDuplicates(charArray) {
			fmt.Printf("Ergebnis: %v", resultIndex)
			break
		}
	}
}

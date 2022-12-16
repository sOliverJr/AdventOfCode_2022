package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func FormatFileInput(FileInput []string) [][]string {
	var FormattedFileInput [][]string
	var ElfPair []string

	for _, element := range FileInput {
		ElfPair = strings.Split(element, ",")
		FormattedFileInput = append(FormattedFileInput, ElfPair)
	}
	return FormattedFileInput
}

func TestIfOneRangeContainsTheOther(SingleElfPair []string) bool {
	var RangeOne []string = strings.Split(SingleElfPair[0], "-")
	var RangeTwo []string = strings.Split(SingleElfPair[1], "-")

	RangeOneLower, _ := strconv.Atoi(RangeOne[0])
	RangeOneHigher, _ := strconv.Atoi(RangeOne[1])
	RangeTwoLower, _ := strconv.Atoi(RangeTwo[0])
	RangeTwoHigher, _ := strconv.Atoi(RangeTwo[1])

	if RangeOneLower <= RangeTwoLower && RangeOneHigher >= RangeTwoHigher {
		// fmt.Printf("Range One Contains Range Two:   Range 1: %v --- Range 2: %v \n", RangeOne, RangeTwo)
		return true
	} else if RangeTwoLower <= RangeOneLower && RangeTwoHigher >= RangeOneHigher {
		// fmt.Printf("Range Two Contains Range One:   Range 1: %v -- Range 2: %v \n", RangeOne, RangeTwo)
		return true
	} else {
		return false
	}
}

func TestIfRangesOverlapAtAll(SingleElfPair []string) bool {
	var RangeOne []string = strings.Split(SingleElfPair[0], "-")
	var RangeTwo []string = strings.Split(SingleElfPair[1], "-")

	RangeOneLower, _ := strconv.Atoi(RangeOne[0])
	RangeOneHigher, _ := strconv.Atoi(RangeOne[1])
	RangeTwoLower, _ := strconv.Atoi(RangeTwo[0])
	RangeTwoHigher, _ := strconv.Atoi(RangeTwo[1])

	if RangeOneHigher >= RangeTwoLower && RangeTwoHigher >= RangeOneLower {
		fmt.Printf("Ranges overlap:   Range 1: %v --- Range 2: %v \n", RangeOne, RangeTwo)
		return true
	} else {
		return false
	}
}

func main() {
	var FileInput []string = ReturnInputFromFile("day04_input.txt")
	var FormattedFileInput [][]string = FormatFileInput(FileInput)
	var AmountElfPairsContain int = 0
	var AmountElfPairsOverlap int = 0

	for _, SingleElfPair := range FormattedFileInput {
		if TestIfOneRangeContainsTheOther(SingleElfPair) {
			AmountElfPairsContain += 1
		}
		if TestIfRangesOverlapAtAll(SingleElfPair) {
			AmountElfPairsOverlap += 1
		}
	}

	fmt.Printf("Part 1: %v \n", AmountElfPairsContain)
	fmt.Printf("Part 2: %v \n", AmountElfPairsOverlap)
}

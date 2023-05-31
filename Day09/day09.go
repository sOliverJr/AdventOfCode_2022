package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Kachel struct {
	visited bool
}

func ReturnInputFromFile(path string) [][]string {
	var InputArray []string
	var ReturnValue [][]string
	f, _ := os.Open(path)

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		InputArray = append(InputArray, scanner.Text())
	}

	for _, line := range InputArray {
		ArrayLine := strings.Split(line, " ")
		ReturnValue = append(ReturnValue, ArrayLine)
	}
	return ReturnValue
}

func FindMinMax(liveCoord []int, minMaxCoord []int) []int {
	if liveCoord[0] > minMaxCoord[0] {
		minMaxCoord[0] = liveCoord[0]
	} else if liveCoord[0] < minMaxCoord[1] {
		minMaxCoord[1] = liveCoord[0]
	} else if liveCoord[1] > minMaxCoord[2] {
		minMaxCoord[2] = liveCoord[1]
	} else if liveCoord[1] < minMaxCoord[3] {
		minMaxCoord[3] = liveCoord[1]
	}
	return minMaxCoord
}

func FindRanges(input [][]string) []int {
	// x, y
	var Coord []int = []int{0, 0}
	// xMax, xMin, yMax, yMin
	var RangeCoord []int = []int{0, 0, 0, 0}

	for _, command := range input {
		var direction string = command[0]
		var amount, _ = strconv.Atoi(command[1])

		switch direction {
		case "U":
			Coord[1] += amount
		case "D":
			Coord[1] -= amount
		case "L":
			Coord[0] -= amount
		case "R":
			Coord[0] += amount
		}

		RangeCoord = FindMinMax(Coord, RangeCoord)
	}

	return RangeCoord
}

func main() {
	var FileInput [][]string = ReturnInputFromFile("Day09/day09_input.txt")
	MinMaxCoords := FindRanges(FileInput)

	fmt.Printf("x_max: %v, x_min: %v, y_max: %v, y_min: %v \n", MinMaxCoords[0], MinMaxCoords[1], MinMaxCoords[2], MinMaxCoords[3])
	fmt.Printf("x: %v, y: %v", MinMaxCoords[0]-MinMaxCoords[1], MinMaxCoords[2]-MinMaxCoords[3])

}

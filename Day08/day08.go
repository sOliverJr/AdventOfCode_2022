// package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReturnInputFromFile(path string) [][]int {
	var InputArray []string
	var ReturnValue [][]int
	var row []int
	f, _ := os.Open(path)

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		InputArray = append(InputArray, scanner.Text())
	}

	for _, line := range InputArray {
		for _, value := range strings.Split(line, "") {
			value, _ := strconv.Atoi(value)
			row = append(row, value)
		}
		ReturnValue = append(ReturnValue, row)
		row = nil
	}
	return ReturnValue
}

func VisibleFromLeftOrRight(grid [][]int, y int, x int, height int) bool {
	var VisibleLeft bool = true
	var VisibleRight bool = true

	if x == 0 || x == len(grid[y])-1 {
		return true
	}
	for i := 0; i != x; i++ {
		if grid[y][i] >= height {
			VisibleLeft = false
			break
		}
	}
	for i := len(grid[y]) - 1; i != x; i-- {
		if grid[y][i] >= height {
			VisibleRight = false
			break
		}
	}

	if VisibleLeft || VisibleRight {
		return true
	} else {
		return false
	}
}

func VisibleFromTopOrBottom(grid [][]int, y int, x int, height int) bool {
	var VisibleTop bool = true
	var VisibleBottom bool = true

	if y == 0 || y == len(grid)-1 {
		return true
	}
	for i := 0; i != y; i++ {
		if grid[i][x] >= height {
			VisibleTop = false
			break
		}
	}
	for i := len(grid[y]) - 1; i != y; i-- {
		if grid[i][x] >= height {
			VisibleBottom = false
			break
		}
	}

	if VisibleTop || VisibleBottom {
		return true
	} else {
		return false
	}
}

func main() {
	var FileInput [][]int = ReturnInputFromFile("Day08/day08_input.txt")
	var AmountVisible int = 0

	for y, line := range FileInput {
		for x, height := range line {
			if VisibleFromTopOrBottom(FileInput, y, x, height) || VisibleFromLeftOrRight(FileInput, y, x, height) {
				AmountVisible++
			}
		}
	}

	fmt.Printf("Result: %v", AmountVisible)
}

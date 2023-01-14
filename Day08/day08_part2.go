package main

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

func VisibleFromLeftOrRight(grid [][]int, y int, x int, height int) (int, int) {
	var VisibleTreesLeft int = 0
	var VisibleTreesRight int = 0

	if x != 0 && x != len(grid[y])-1 {
		for i := x + 1; i < len(grid[y]); i++ {
			if grid[y][i] < height {
				VisibleTreesRight++
			} else if grid[y][i] >= height {
				VisibleTreesRight++
				break
			}
		}

		for i := x - 1; i >= 0; i-- {
			if grid[y][i] < height {
				VisibleTreesLeft++
			} else if grid[y][i] >= height {
				VisibleTreesLeft++
				break
			}
		}
	}

	return VisibleTreesLeft, VisibleTreesRight
}

func VisibleFromTopOrBottom(grid [][]int, y int, x int, height int) (int, int) {
	var VisibleTreesTop int = 0
	var VisibleTreesBottom int = 0

	if y != 0 && y != len(grid)-1 {
		for i := y + 1; i < len(grid); i++ {
			if grid[i][x] < height {
				VisibleTreesBottom++
			} else if grid[i][x] >= height {
				VisibleTreesBottom++
				break
			}
		}

		for i := y - 1; i >= 0; i-- {
			if grid[i][x] < height {
				VisibleTreesTop++
			} else if grid[i][x] >= height {
				VisibleTreesTop++
				break
			}
		}
	}
	return VisibleTreesTop, VisibleTreesBottom
}

func main() {
	var FileInput [][]int = ReturnInputFromFile("Day08/day09_input.txt")
	var HighestScenicScore int = 0

	for y, line := range FileInput {
		for x, height := range line {
			TreesLeft, TreesRight := VisibleFromLeftOrRight(FileInput, y, x, height)
			TreesTop, TreesBottom := VisibleFromTopOrBottom(FileInput, y, x, height)
			var ScenicScore int = TreesTop * TreesBottom * TreesLeft * TreesRight
			if ScenicScore > HighestScenicScore {
				HighestScenicScore = ScenicScore
			}
		}
	}

	fmt.Printf("Result: %v", HighestScenicScore)
}

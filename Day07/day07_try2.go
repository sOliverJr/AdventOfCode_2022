package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReturnInputFromFileAsSLices(path string) [][]string {
	var InputArray [][]string
	f, _ := os.Open(path)

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		InputArray = append(InputArray, strings.Split(scanner.Text(), " "))
	}
	return InputArray
}

func CalculateDirSize(DirName string, StartLineNumber int, InputFile [][]string) int {
	var DirSize int = 0
	var ReachedCurrentDir bool = false
	for CurrentLineNumber, CurrentLine := range InputFile {
		if CurrentLineNumber < StartLineNumber {
			continue
		}
		// End of dir ls is reached
		if ReachedCurrentDir && CurrentLine[1] == "cd" {
			ReachedCurrentDir = false
			if DirSize <= 100000 {
				result += DirSize
			}
			return DirSize
		}
		// Add File size to dir size
		if ReachedCurrentDir && CurrentLine[0] != "dir" && CurrentLine[1] != "ls" {
			size, _ := strconv.Atoi(CurrentLine[0])
			DirSize += size
		}
		// Calculate subfolder size
		if ReachedCurrentDir && CurrentLine[0] == "dir" {
			DirSize += CalculateDirSize(CurrentLine[1], CurrentLineNumber, InputFile)
		}
		// Directory is reached
		if CurrentLine[0] == "$" && CurrentLine[1] == "cd" && CurrentLine[2] == DirName {
			ReachedCurrentDir = true
		}
	}
	return 1000000
}

var result int = 0

func main() {
	var InputFile = ReturnInputFromFileAsSLices("/Users/oli/Documents/GitHub/AdventOfCode_2022/Day07/day07_input.txt")
	var RootDirSize int = CalculateDirSize("/", 0, InputFile)
	fmt.Println("Result: ", result)
	fmt.Println("Root Directory Size: ", RootDirSize)
}

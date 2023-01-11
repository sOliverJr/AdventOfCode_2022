package main

import (
	"bufio"
	"os"
	"strings"
)

type file struct {
	name string
	size int
}

type directory struct {
	name        string
	directories []directory
	files       []file
	size        int
}

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

func lineIsCommand(line string) (bool, []string) {
	var lineSlice []string = strings.Split(line, " ")
	if line[0] == '$' {
		return true, lineSlice
	} else {
		return false, lineSlice
	}
}

func calcSingleDirSize(dir directory) int {
	var size int = 0
	for _, subDir := range dir.directories {
		size += subDir.size
	}
	for _, file := range dir.files {
		size += file.size
	}
	return size
}

func subDirsSizeIsCalculated(dir directory) bool {
	for _, subDir := range dir.directories {
		if subDir.size == 0 {
			return false
		}
	}
	return true
}

func calcAllDirSizes(rootDir directory) directory {
	for i, subDir := range rootDir.directories {
		if subDirsSizeIsCalculated(subDir) || len(subDir.directories) == 0 {
			rootDir.directories[i].size = calcSingleDirSize(subDir)
		} else {
			rootDir.directories[i] = calcAllDirSizes(subDir)
		}
	}
	rootDir.size = calcSingleDirSize(rootDir)
	return rootDir
}

func sumAllDirsUnder100000B(rootDir directory) int {
	var sum int = 0
	for _, subDir := range rootDir.directories {
		sum += sumAllDirsUnder100000B(subDir)
	}
	if rootDir.size <= 100000 {
		sum += rootDir.size
	}
	return sum
}

func popDirectory(dirSlice []directory) (directory, []directory) {
	return dirSlice[len(dirSlice)-1], dirSlice[:len(dirSlice)-1]
}

func main() {
	var FileInput []string = ReturnInputFromFile("day07_input.txt")
	var dirs []directory
	var files []file

	var rootFolder directory = directory{
		name:        "/",
		directories: dirs,
		files:       files,
		size:        0,
	}
	// var currentDirectory []directory = []directory{rootFolder}

	var dirContentIsBeeingListed bool = false

	for _, line := range FileInput {
		lineIsCommand, commandSlice := lineIsCommand(line)
		if lineIsCommand {
			if commandSlice[1] == "ls" {
				dirContentIsBeeingListed = true
			} else {
				dirContentIsBeeingListed = false
			}
		} else if !dirContentIsBeeingListed {

		}
	}

	rootFolder = calcAllDirSizes(rootFolder)

}

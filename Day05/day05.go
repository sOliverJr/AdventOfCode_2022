package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rearrangement struct {
	amount int
	from   int
	to     int
}

type stack struct {
	number  int
	content []string
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

func FormatFileInput(fileInput []string) []rearrangement {
	var rearrangementArray []rearrangement

	for _, operation := range fileInput {
		var operationArray []string = strings.Split(operation, " ")

		var amount, _ = strconv.Atoi(operationArray[1])
		var from, _ = strconv.Atoi(operationArray[3])
		var to, _ = strconv.Atoi(operationArray[5])

		anweisung := rearrangement{
			amount: amount,
			from:   from,
			to:     to,
		}

		rearrangementArray = append(rearrangementArray, anweisung)
	}
	return rearrangementArray
}

func popStack(popStack stack) (string, stack) {
	crate, stackContent := popStack.content[len(popStack.content)-1], popStack.content[:len(popStack.content)-1]
	return crate, stack{
		number:  popStack.number,
		content: stackContent,
	}
}

func appendStack(appendStack stack, element string) stack {
	var content []string = appendStack.content
	content = append(content, element)
	appendStack.content = content
	return appendStack
}

func executeAnweisung(currentArrangement []stack, anweisung rearrangement) []stack {
	for i := 0; i < anweisung.amount; i++ {
		crate, newStackRemoved := popStack(currentArrangement[anweisung.from-1])
		currentArrangement[anweisung.from-1] = newStackRemoved

		newStackAdded := appendStack(currentArrangement[anweisung.to-1], crate)
		currentArrangement[anweisung.to-1] = newStackAdded
	}

	return currentArrangement
}

func executeAnweisungPartTwo(currentArrangement []stack, anweisung rearrangement) []stack {
	var crateStack []string

	for i := 0; i < anweisung.amount; i++ {
		crate, newStackRemoved := popStack(currentArrangement[anweisung.from-1])
		currentArrangement[anweisung.from-1] = newStackRemoved
		crateStack = append(crateStack, crate)
	}
	// Reverse-iterate trough slice
	for i := range crateStack {
		currentArrangement[anweisung.to-1] = appendStack(currentArrangement[anweisung.to-1], crateStack[len(crateStack)-1-i])
	}

	return currentArrangement
}

func main() {
	var result string = ""
	var currentArrangement []stack = []stack{
		{number: 1, content: []string{"F", "D", "B", "Z", "T", "J", "R", "N"}},
		{number: 2, content: []string{"R", "S", "N", "J", "H"}},
		{number: 3, content: []string{"C", "R", "N", "J", "G", "Z", "F", "Q"}},
		{number: 4, content: []string{"F", "V", "N", "G", "R", "T", "Q"}},
		{number: 5, content: []string{"L", "T", "Q", "F"}},
		{number: 6, content: []string{"Q", "C", "W", "Z", "B", "R", "G", "N"}},
		{number: 7, content: []string{"F", "C", "L", "S", "N", "H", "M"}},
		{number: 8, content: []string{"D", "N", "Q", "M", "T", "J"}},
		{number: 9, content: []string{"P", "G", "S"}},
	}
	var FileInput []string = ReturnInputFromFile("day05_input.txt")
	var rearrangementsArray []rearrangement = FormatFileInput(FileInput)

	for _, anweisung := range rearrangementsArray {
		currentArrangement = executeAnweisungPartTwo(currentArrangement, anweisung)
	}

	for _, stack := range currentArrangement {
		crate, _ := popStack(stack)
		result += crate
	}

	fmt.Printf("Teil 1: %v", result)
}

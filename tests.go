package main

import "fmt"

func main() {
	var array []string = []string{"test", "hallo"}

	if array[0] == "hey" && array[2] == "ups" {
		fmt.Print("Ups")
	} else {
		fmt.Print("Success")
	}
}

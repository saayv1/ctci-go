package main

import (
	"fmt"
	"strings"
)

func isRotation(input1 string, input2 string) bool {
	if len(input1) != len(input2) {
		return false
	}
	newStr := input2 + input2

	return strings.Contains(newStr, input1)
}

func main() {
	input1 := "waterbottle"
	input2 := "erbottlewat"
	fmt.Println(isRotation(input1, input2))
}

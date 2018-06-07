package main

import (
	"strconv"
	//	"strings"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "aabcccccaaa"
	var compressed string
	firstAlphabet := string(input[0])
	count := 0

	for _, letters := range input {
		if string(letters) == firstAlphabet {
			count++
		} else {

			compressed = compressed + firstAlphabet + strconv.Itoa(count)
			firstAlphabet = string(letters)
			count = 1
		}
	}
	compressed = compressed + firstAlphabet + strconv.Itoa(count)
	l1 := utf8.RuneCountInString(input)
	l2 := utf8.RuneCountInString(compressed)
	if l2 > l1 {
		fmt.Println(input)
		return
	}
	fmt.Println(compressed)
}

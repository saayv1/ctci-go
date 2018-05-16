/*
*
* @author: Vyaas Narendra Shenoy
*
* Given a string write a function to check if it is a permutation of a palindrome. A palindrome is a word or phrase that os the same forwards and backwards. A permutation is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words
*
*Example
*
*Input  Taco cat
*
*Output True
*
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	input, err := in.ReadString('\n')

	m := make(map[string]int)

	for _, runes := range input {
		x, ok := m[string(runes)]
		if ok {
			x++
			m[strings.ToLower(string(runes))] = x
		} else {
			if string(runes) == " " || string(runes) == "\n" {
			} else {
				m[strings.ToLower(string(runes))] = 1
			}
		}
		fmt.Println(string(runes))
	}

	oddCount := 0
	for key, value := range m {
		if value%2 != 0 {
			oddCount++
		}
		fmt.Println("Key:", key, "Value:", value)
	}

	if oddCount != 1 {
		fmt.Println("False")
		return
	}

	fmt.Println("True")
	if err != nil {
		panic(err)
	}

}

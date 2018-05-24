package main

import (
	"fmt"
//	"unicode/utf8"
)

func main() {
	w1 := "pale"
	w2 := "pae;"

//	l1 := utf8.RuneCountInString(w1)
//	l2 := utf8.RuneCountInString(w2)

	m1 := make(map[string]int)
	m2 := make(map[string]int)

	for _, runes := range w1 {
		x, ok := m1[string(runes)]
		if ok {
			x++
			m1[string(runes)] = x
		} else {
			m1[string(runes)] = 1
		}
	}

	for _, runes := range w2 {
		x, ok := m2[string(runes)]
		if ok {
			x++
			m2[string(runes)] = x
		} else {
			m2[string(runes)] = 1
		}
	}

	mismatch := 0
	for _, runes := range w1 {
		_, ok := m2[string(runes)]
		if !ok {
			mismatch++
		}
	}
	if mismatch > 1 {
		fmt.Println("FALSE")
		return
	} else {
		fmt.Println("TRUE")
		return
	}

}

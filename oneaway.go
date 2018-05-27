/*
*
* @author: Vyaas Narendra Shenoy
*
* There are 3 type of edits that can be performed on strings: insert a character, remove a character or replace a character. Given 2 strings write a function to check if they are 1 edit or 0 edits away
*Example
*
*Input  pale,ple
*
*Output True
*
 */

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	w1 := "pale"
	w2 := "bake"
	/*
	   check length first
	   if they are same then check for difference individually

	   if they are not the same
	   then check if the difference in lenght
	*/

	l1 := utf8.RuneCountInString(w1)
	l2 := utf8.RuneCountInString(w2)

	if l1 == l2 {
		var dissimilarity int
		dissimilarity = 0
		for i, runes := range w1 {
			if string(runes) != string(w2[i]) {
				dissimilarity++
			}
			if dissimilarity > 1 {
				fmt.Println("FALSE")
				return
			}
		}
	}	else	{
			if l1>l2 {
				dissimilarity := 0
				for i,j:=0,0;i<l1;i,j=i+1,j+1 {
					if string(w1[i])!= string(w2[j]) {
						i++;
						dissimilarity++;		
					}
					if dissimilarity >1 {
						fmt.Println("FALSE")
						return
					}
				}
			}
 			if l2>l1 {
                                dissimilarity := 0
                                for i,j:=0,0;i<l2;i,j=i+1,j+1 {
                                        if string(w1[j])!= string(w2[i]) {
                                                i++;    
                                                dissimilarity++;
                                        }
                                        if dissimilarity >1 {
                                                fmt.Println("FALSE")
                                                return
                                        }
         			} 
			}                
	}
	fmt.Println("TRUE")
	/*
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
	*/
}

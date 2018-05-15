package main

import (
	"fmt"
	//"math"
	"strings"
	"os"
	"bufio"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s1 string
	
	var s2[] string
	
	s1,err := in.ReadString('\n');
	_=err;

	for _,runes := range s1 {
		
		if string(runes) == " "{
			s2 = append(s2,"%20");
		}else{
		s2 = append(s2,string(runes));
		}
	}	
	
	s1 =strings.Join(s2,"")
	fmt.Println(s1);
}	

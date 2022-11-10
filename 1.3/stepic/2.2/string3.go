package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	var a string
	//fmt.Println(strings.Trim(s, "zab"))
	RS := []rune(s)
	for i := range RS {
		fmt.Println(i)
		if strings.Count(s, string(RS[i])) > 1 {
			continue
		} else {
			//var a string
			a += string(RS[i])
		}
		//fmt.Println(a)
		//b += strings(RS[i])
		//b:=strings.ReplaceAll()
	}
	fmt.Println(a)
}

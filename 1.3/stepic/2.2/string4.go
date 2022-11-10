package main

import (
	"fmt"
	"unicode"
)

func c(s string) bool {

	RS := []rune(s)
	lenght := len(RS)
	var d bool
	if lenght > 5 {
		d = false
	} else {

	}
	for _, i := range RS {

		//fmt.Println(lenght)

		//if unicode.IsDigit(RS[len(i)])>5{
		//continue
		if (unicode.Is(unicode.Latin, i)) && (unicode.Is(unicode.Digit, i)) {
			d = false
			//fmt.Println(unicode.Is(unicode.Latin, i))
			break

		} else {
			d = true
		}

	}
	return d
}

func main() {
	var s string
	fmt.Scan(&s)
	//var c bool
	if c(s) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}

}

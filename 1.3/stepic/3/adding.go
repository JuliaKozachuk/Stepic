package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var a string
	var b string
	fmt.Scan(&a, &b)
	fmt.Println(adding(a, b))

}
func clean(x string) int64 {
	//var x string
	//fmt.Scan(&x)
	resX := []rune(x)

	for _, i := range resX {
		//fmt.Println(i)
		if !unicode.IsDigit(i) {
			x = strings.ReplaceAll(x, string(i), "")
			//fmt.Println(x)
		}
	}
	res, _ := strconv.ParseInt(x, 10, 64)
	fmt.Println(res)
	return res

}
func adding(y, z string) int64 {
	onenumber := clean(y)
	twonumber := clean(z)
	return onenumber + twonumber
}

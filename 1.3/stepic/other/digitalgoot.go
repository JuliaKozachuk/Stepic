package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	var digitalroot int

	// onenumber := number / 1000
	// //fmt.Println(onenumber)
	// twonumber := number % 1000 / 100

	// //fmt.Println(twonumber)
	// threenumber := number / 10 % 10
	// //fmt.Println(threenumber)
	// fournumber := number % 10
	// //fmt.Println(fournumber)

	// root := onenumber + twonumber + threenumber + fournumber
	// //fmt.Println(root)

	// oneroot := root % 100 / 10
	// //fmt.Println(oneroot)
	// tworoot := root % 10
	// //fmt.Println(tworoot)
	// digitalroot = oneroot + tworoot
	// fmt.Println(digitalroot)

	//for i:=0;i<10;i++{
	if number > 10 {
		digitalroot = 1 + (number-1)%9

	}
	fmt.Println(digitalroot)

}

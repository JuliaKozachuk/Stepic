package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)

	b := a % 10
	//fmt.Println(b)
	c := (a % 100) / 10
	//fmt.Println(c)
	d := a / 100
	//fmt.Println(d)
	e := b + c + d
	fmt.Println(e)

}

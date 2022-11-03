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

	//fmt.Println(b, c, d)
	//fmt.Print(b)
	//fmt.Print(c)
	//fmt.Print(d)
	//fmt.Println(e)
	fmt.Printf("%d%d%d", b, c, d)
}

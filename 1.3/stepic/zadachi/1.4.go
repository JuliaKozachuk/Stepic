package main

import "fmt"

func main() {
	var a float32
	fmt.Scan(&a)
	var b float32
	fmt.Scan(&b)
	var c float32

	c = (a + b) / 2
	fmt.Println(c)

}

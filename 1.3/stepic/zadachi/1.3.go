package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	var c int
	fmt.Scan(&c)

	if a+b > c || a+c > b || b+c > a {
		fmt.Println("Cуществует")
	}
	if a+b < c || a+c < b || b+c < a {
		fmt.Println("Не существут")
	}

}

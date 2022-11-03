package main

import "fmt"

func main() {

	var a int
	var b int
	var c []int
	fmt.Scan(&a)

	for i := 0; i <= a; i++ {
		fmt.Scan(&b)
		c[i-1] = b
		fmt.Println(b)
	}
}

package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)
	var b int

	var c int
	for ; a < 10; a++ {
		fmt.Scan(&b)

		if 10 <= b && b%8 == 0 && b <= 100 {

			c += b

		}

	}
	fmt.Println(c)
}

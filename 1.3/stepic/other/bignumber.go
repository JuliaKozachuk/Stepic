package main

import "fmt"

func main() {
	var a, b int
	var c int
	var d int
	fmt.Scan(&a, &b)

	for i := b; i >= a; i-- {

		if i%7 == 0 {
			c = i

			//fmt.Println(c)

			d++

			break

		} else {
			continue
		}

	}
	if d == 0 {
		fmt.Println("NO")
	}
	fmt.Println(c)

}

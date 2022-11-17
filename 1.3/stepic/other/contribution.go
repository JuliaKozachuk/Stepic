package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	var p int
	fmt.Scan(&p)
	var y int
	fmt.Scan(&y)
	years := 0

	// for i := 0; i > y; i++ {
	// 	if x < y {
	// 		years = x * p / 100
	// 		continue
	// 		//} else {

	// 	} else {
	// 		fmt.Println(years)

	// 	}
	// }
	for {
		s := x * p / 100
		x += s
		years++
		if x > y {
			break

		}

	}
	fmt.Println(years)
}

package main

import "fmt"

func main() {

	var (
		a int
		b int
		c int
		d int
	)
	fmt.Scan(&a, &b, &c, &d)

	//minimumFromFour(a, b, c, d)
	// // fmt.Scan(&a)
	// // fmt.Scan(&b)
	// // fmt.Scan(&c)
	// // fmt.Scan(&d)
	fmt.Println(minimumFromFour(a, b, c, d))
}

func minimumFromFour(a int, b int, c int, d int) int {
	var min int

	if a < b && a < c && a < d && a == a {
		min = a
	} else if b < c && b < a && b < d && b == b {
		min = b
	} else if c < d && c < a && c < b && c == c {
		min = c
	} else if d < a && d < b && d < c && d == d {
		min = d
	}
	return min

}

// func minimumFromFour(a int, b int, c int, d int) int {
// 	fmt.Scan(&a, &b, &c, &d)
// 	fmt.Println(a)
// 	// fmt.Scan(&a)
// 	// fmt.Scan(&b)
// 	// fmt.Scan(&c)
// 	// fmt.Scan(&d)
// 	var min int
// 	//for i := 0;i<10:i++:{

// 	if a < b {
// 		min = a
// 	} else if b < c {
// 		min = b
// 	} else if c < d {
// 		min = c
// 	} else if d < a {
// 		min = c
// 	}
// 	return min
// }

// if e < b && a < c && a < d && a >= 0{
// fmt.Println(a)

// }

// fmt.Println(e)

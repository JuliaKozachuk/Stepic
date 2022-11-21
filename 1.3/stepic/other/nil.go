package main

import "fmt"

func main() {
	var array []int
	var N int
	var c int
	fmt.Scan(&N)
	//min := 0
	array = make([]int, N)
	//fmt.Println(array)

	for i := range array {

		fmt.Scan(&array[i])

		array[i] = c
		//fmt.Println(c)

		if c == 0 {

			c++
			//c = min

			// }
		}
		//fmt.Println(array)

	}
	fmt.Println(c)
}

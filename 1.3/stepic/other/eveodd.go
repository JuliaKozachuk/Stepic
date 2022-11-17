package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	array := make([]int, N)
	//fmt.Println(array)

	// for i := 0; i < n; i++ {
	// 	var s int
	// 	fmt.Scan(&s)
	// 	array[i] = s
	//fmt.Println(s)

	//fmt.Println(array[3])

	for d := range array {
		fmt.Scan(&array[d])
		if d%2 == 0 {
			fmt.Println("\n", array[d], " ")

		}

		//} else {
		//fmt.Println(d)

		//}

	}

}

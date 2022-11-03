package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	var a int

	massiv := make([]int, N)

	for i := 0; i < N; i++ {
		//fmt.Scan(&massiv[i])
		fmt.Scan(&a)
		massiv[i] = a

		//fmt.Println(massiv)

	}
	fmt.Println(massiv)
	max := massiv[0]
	for i := 0; i > max; i++ {
		if massiv[i]%max == 0 {
			max = massiv[i]
			//}
		}
		fmt.Println(max)
	}
}

package main

import "fmt"

func main() {
	// 	var a uint8
	// 	fmt.Scan(&a)
	// 	//var b uint8
	// 	var i uint8
	// 	var workArray []uint8

	// 	for i = 0; i > 100; i = i + a {
	// 		workArray = append(workArray, i)

	// 	}
	// 	fmt.Println(workArray)

	// }
	var workArray [10]uint8

	//var b uint8
	//var r uint8
	for i := 0; i < 10; i++ {
		var a uint8
		fmt.Scan(&a)
		workArray[i] = a

	}
	for j := 0; j < 3; j++ {
		var i1, i2 uint8
		fmt.Scan(&i1, &i2)
		workArray[i1], workArray[i2] = workArray[i2], workArray[i1]
	}
	for _, s := range workArray {
		fmt.Println(s, " ")

	}

}

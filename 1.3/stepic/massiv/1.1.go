package main

import (
	"fmt"
)

func main() {
	//var a uint8
	//var b uint8
	//var c uint8
	//var d uint8
	var workArray [10]uint8

	// fmt.Println(workArray)
	// for i := 0; i < len(workArray); i++ {
	// 	//fmt.Println(workArray[i])

	// 	fmt.Scan(&workArray[i])

	//fmt.Scan(workArray&a)
	//workArray[i] = a

	//}
	// for i := 0; i <= 3; i++ {

	// 	fmt.Scan(&b, &c)
	// 	d = workArray[b]
	// 	workArray[b] = workArray[c]
	// 	workArray[c] = d
	// }
	// fmt.Println(workArray)
	for a := 0; a < len(workArray); a++ {
		fmt.Scan(&workArray[a])
	}
	fmt.Println(workArray)
	for _, elem := range workArray {
		fmt.Println(elem)

	}

	// for a, b := range workArray {
	// fmt.Println(a, b)
}

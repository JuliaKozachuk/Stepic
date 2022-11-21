package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var c int
	var min int
	//min = 0
	var d int
	var array []int
	array = make([]int, N)

	for i := range array {
		fmt.Scan(&c)
		array[i] = c

	}
	min = array[0]
	for j := 0; j < len(array); j++ {
		if min > array[j] {
			min = array[j]
			d = 1
		} else if min == array[j] {
			d++
		}
	}
	fmt.Println(d)

}

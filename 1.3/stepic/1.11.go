package main

import "fmt"

func main() {

	var R float64
	fmt.Scan(&R)
	//R = (R * 2)
	//fmt.Printf("%2.2f", R)

	switch {
	case R <= 0:
		fmt.Printf("число %2.2f не подходит", R)

	case R >= 10000:
		fmt.Printf("%e", R)
	default:
		var x float64 = R * R
		fmt.Printf("%.4f", x)
	}
}

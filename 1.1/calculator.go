package main

import "fmt"

func main() {

	for a := 1; a <= 10; a++ {

		for b := 1; b <= 10; b++ {
			//fmt.Print(a*b, "\t")
			fmt.Print(a, " * ", b, " =", a*b)
			fmt.Println()
		}
		//fmt.Println()
	}

}

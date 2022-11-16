package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	c := 0

	for i := 0; i < a; i++ {
		//c = i
		var b int
		//var c int
		//c = 0
		fmt.Scan(&b)
		//c := a + b

		if b < 100 && b >= 10 && b%8 == 0 {
			c += b

		}
		//if b%8==0 && b<100{

	}
	fmt.Println(c)

}

//fmt.Println(c)
//}

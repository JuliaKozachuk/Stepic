package main

import "fmt"

func test(x1 *int, x2 *int) {
	//*x1, *x2 = 2, 2
	z := *x1 * *x2
	// //z := *x1 * *x2
	fmt.Println(z)

}
func main() {

	x1, x2 := 2, 2
	//fmt.Scan(&x1, &x2)
	test(&x1, &x2)
	// 	z := x1 * x2
	fmt.Println(x1 * x2)
	//func main(){

	//z:=x1*x2

}

//}

// func test(x1 *int, x2 *int) {
// 	*x1 = 2
// 	*x2 = 2
// 	z := *x1 * *x2
// 	//z := *x1 * *x2
// 	fmt.Println(z)

// }

package main

import "fmt"

func test(x1 *int, x2 *int) {
	//func test(x1 *int, x2 *int) {
	//z := *x1
	//y := *x2
	//var y int
	//var z int
	//*x1, *x2 = *x2, *x1
	y := *x1
	*x1 = *x2
	*x2 = y

	fmt.Println(*x1, *x2)

}
func main() {
	x1, x2 := 2, 4
	test(&x1, &x2)
}

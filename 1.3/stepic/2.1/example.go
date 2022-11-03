package main

import "fmt"

// func main() {
// 	add(4, 5)  // x + y = 9
// 	add(20, 6) // x + y = 26
// }

//	func add(x int, y int) {
//		var z = x + y
//		fmt.Println("x + y = ", z)
//	}

func fn() (int, error) {
	// Какая-то полезная работа
	// ...
	return 0, nil
}
func ExampleIgnor() {
	fn()

	i, _ := fn()
	fmt.Println(i)

	_, err := fn()
	if err == nil {
		fmt.Println("Ошибок нет")
	}

	// Output:
	// 0
	// Ошибок нет
}

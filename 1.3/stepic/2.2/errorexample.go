package main

import "fmt"

func divide(a int, b int) int {
	return a / b
}

func main() {
	var input int
	fmt.Scan(&input)
	fmt.Println(divide(input, 5)) //Выведем результат
}

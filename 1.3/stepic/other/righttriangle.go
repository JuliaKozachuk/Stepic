package main

import "fmt"

func main() {
	var a int

	var b int

	var c int
	fmt.Scan(&a, &b, &c)
	// if a<c{
	// 	sideA:=a
	// }
	// if b<c{
	// 	sideB:=b
	// }
	// var sideC int
	// var sideA int
	// var sideB int

	// if c > a && c > b {
	// 	sideC = c
	// 	sideA = a
	// 	sideB = b
	// 	fmt.Println(sideA, sideB, sideC)
	// } else if a > c && a > b {
	// 	sideC = a
	// 	sideA = c
	// 	sideB = b
	// 	fmt.Println(sideA, sideB, sideC)
	// } else {
	// 	sideC = b
	// 	sideA = a
	// 	sideB = c
	// 	fmt.Println(sideA, sideB, sideC)
	// }
	// if (sideA*sideA)+(sideB+sideB) == sideC*sideC {
	// 	fmt.Println("Прямоугольный")

	// } else {
	// 	fmt.Println("Не прямоугольный")
	// }

	if (a*a)+(b*b) == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
	// if (c*c)+(a*a) == b*b {
	// 	fmt.Println("Прямоугольный")
	// } else {
	// 	fmt.Println("Не прямоугольный")
	// }
	// if (c*c)+(b*b) == a*a {
	// 	fmt.Println("Прямоугольный")
	// } else {
	// 	fmt.Println("Непрямоугольный")
	// }

}

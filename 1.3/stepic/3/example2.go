package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 	s := "23.23456"
	// 	// как и в прошлом шаге, здесь 2 параметр - bitSize
	// 	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	// 	// но нужно понимать что метод все равно вернет float64
	// 	result, err := strconv.ParseFloat(s, 64)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(result)         // 23.23456
	// 	fmt.Printf("%T \n", result) // float64

	//		// Конкретный пример для разных bitSize:
	//		s = "1.0000000012345678"
	//		//  не будем обрабатывать ошибки в примерах, но на практике так не делайте ;)
	//		result32, _ := strconv.ParseFloat(s, 32)
	//		result64, _ := strconv.ParseFloat(s, 64)
	//		fmt.Println("bitSize32:", result32) // вывод 1 (не уместились)
	//		fmt.Println("bitSize64:", result64) // вывод  1.0000000012345678
	//	}
	//res := []rune("stepik")
	//res := string([]byte("stepik"))
	//res := 101.0 / 10
	//a := 10.123
	//res := int64(a)
	//res := strconv.Itoa(int(float64(100/10)) + 1)
	//res := strconv.FormatBool(10 == int16(float64(100/10)))
	res := strconv.FormatBool(10.1 == float32(float64(100/10.1)))
	fmt.Println(res)
}

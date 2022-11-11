package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		err := errors.New("на ноль делить нельзя")
		return 0, err
	}
	return a / b, nil
}
func main() {
	var a int
	var b int
	//_, err := fmt.Scan(&a_input)
	//_, off := fmt.Scan(&b_input)
	fmt.Scan(&a, &b)
	//if err != nil && off != nil {
	//fmt.Println("ошибка")
	//} else {
	//int, error := (divide(a_input, b_input))
	//errors := nil

	//fmt.Println(int, error) //Выведем результат, если ошибок нет
	//}
	//}
	res, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
		return
	}
	fmt.Println(res)
}

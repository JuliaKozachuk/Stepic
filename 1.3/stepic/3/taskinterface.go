package main

import (
	"fmt"
	// пакет используется для проверки ответа, не удаляйте его
	// пакет используется для проверки ответа, не удаляйте его
)

// func readTask(value1, value2, operation interface{}) {
// 	// switch v:=v.(type){
// 	// case "+":
// 	// }
// 	return 8.0, 5.6, "*"
// }
// func Messageerror(value interface{}) {
// 	fmt.Printf("value=%v:%T\n", value, value)
// }

func main() {
	value1, value2, operation := readTask()
	//value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
	v1, ok := value1.(float64)
	if !ok {
		Messageerror(value1)
		return
	}
	v2, ok := value2.(float64)
	if !ok {
		Messageerror(value2)
		return

	}
	if oper, ok := operation.(string); ok {
		switch oper {
		case "+":
			result := v1 + v2
			fmt.Printf("%.4f\n", result)
		case "-":
			result := v1 - v2
			fmt.Printf("%.4f\n", result)
		case "*":
			result := v1 * v2
			fmt.Printf("%.4f\n", result)
		case "/":
			result := v1 / v2
			fmt.Printf("%.4f\n", result)
		default:
			fmt.Printf("Неизвестная операция")
			return

		}
	} else {
		fmt.Println("Неизвестная операция")
		return
	}
}
func readTask() (value1, value2, operation interface{}) {
	// switch v:=v.(type){
	// case "+":
	// }
	return 12.2, 5.6, "+"
}

func Messageerror(value interface{}) {
	fmt.Printf("value=%v: %T\n", value, value)
}

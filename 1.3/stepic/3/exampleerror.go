package main

import (
	"fmt"
)

//func main() {

//}

// type customError uint

// func (c customError) Error() string {
// 	return fmt.Sprintf("цифра, индекс %d", c)
// }

// func errorInString(str string) error {
// 	// Полезная работа со строкой проигнорирована
// 	for i, s := range str {
// 		if unicode.IsDigit(s) {
// 			return customError(i)
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	err := errorInString("string1string")
// 	if err != nil {
// 		fmt.Printf("Ошибка обработана: %v\n", err)
// 	}
// 	if cError, ok := err.(customError); ok {
// 		fmt.Printf("Контекст: %d\n", cError)
// 	}

//		// Output:
//		// Ошибка обработана: цифра, индекс 6
//		// Контекст: 6
//	}
type Animal struct {
	Name string
	Age  uint
}

func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func main() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)

	// Output: Gopher (2)
}

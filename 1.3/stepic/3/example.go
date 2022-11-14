// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := "str"

// 	b := []byte(a)

// 	c := string(b)

// 	fmt.Println(a) // str

// 	fmt.Println(b) // [115 116 114] - побайтовый срез

// 	fmt.Println(c) // str
// }
//Тоже самое работает и со срезами типа rune:

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := "строка"

// 	b := []rune(a) // срез рун

// 	c := string(b)

// 	fmt.Println(a) // строка

// 	fmt.Println(b) // [1089 1090 1088 1086 1082 1072] - срез рун

//		fmt.Println(c) // строка
//	}
// package main

// import "fmt"

//	func main() {
//		a := "12"
//		c := 100
//		b := []byte(a)
//		a += string(b)
//		fmt.Println(c + b)
//	}
package main

import (
	"fmt"
	"strconv"
)

func main() {
	user := "ученик"
	steps := 4

	//fmt.Println("Поздравляю, " + user + "! Ты прошел " + steps + " шага по приведению типов.")
	fmt.Println("Поздравляю," + user + "!Ты прошел" + " " + strconv.Itoa(steps) + " " + "шага по приведению типов.")
}

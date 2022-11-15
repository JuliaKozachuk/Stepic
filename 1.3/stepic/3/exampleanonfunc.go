package main

import (
	"fmt"
)

// func invert(r rune) rune {
// 	// Если буква строчная, то она возвращается заглавной
// 	if unicode.IsLower(r) {
// 		return unicode.ToUpper(r)
// 	}
// 	// Иначе возвращается строчной
// 	return unicode.ToLower(r)
// }
// func main() {
// 	src := "aBcDeFg"
// 	test := "AbCdEfG"

// 	// Обратите внимание, что скобки после имени функции используются только при ее вызове
// 	src = strings.Map(invert, src)

// 	fmt.Printf("Инвертированная строка: %s. Результат: %v.\n", src, src == test)

// 	// Output:
// 	// Инвертированная строка: AbCdEfG. Результат: true.
// }
// func returnFunction() func(rune) rune {
// 	return invert
// }

// //func main() {

// }
// func main() {
// 	src := "aBcDeFg"
// 	test := "AbCdEfG"

// 	// Обратите внимание, что скобки после имени функции используются только при ее вызове
// 	src = strings.Map(func(r rune) rune {
// 		if unicode.IsLower(r) {
// 			return unicode.ToUpper(r)
// 		}
// 		return unicode.ToLower(r)
// 	}, src)

// 	fmt.Printf("Инвертированная строка: %s. Результат: %v.\n", src, src == test)

//		// Output:
//		// Инвертированная строка: AbCdEfG. Результат: true.
//	}
func main() {
	// 	// Присваиваем переменной значение анонимной функции
	// 	fn := func(a, b int) int { return a + b }

	// 	// Выполняем анонимную функцию на месте
	// 	// Обратите внимание на использование скобок при вызове функции
	// 	func(a, b int) {
	// 		fmt.Println(a + b)
	// 	}(12, 34)

	// 	fmt.Println(fn(17, 15))

	// 	// Output:
	// 	// 46
	// 	// 32
	// }

	// В примере мы присвоили переменной fn функцию вида func(int, int) int,
	// затем выполнили другую анонимную функцию, а затем выполнили функцию,
	// присвоенную переменной fn. Обратите внимание на использование скобок в примерах - вызов функции требует наличия скобок,
	// в которых указываются передаваемые функции аргументы (если аргументы не передаются - скобки пустые).
	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
	fmt.Printf("%T", x)
}

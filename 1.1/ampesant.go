//package main
//
//import "fmt"

//func main() {
//	a := 1
//	x := &a
//	*x = 2

//		fmt.Println(a)
//	}
//package main

//import "fmt"

// func main() {
// var a int
// fmt.Println(&a)
// }
//package main

//import "fmt"

//func main() {

//var u *int

//*u = 10

//fmt.Println(u)

// }// выдаст ошибку потому что показетелю u  не выделено место памяти.
//package main

//import "fmt"

// func main() {
// var u *int
// y := 100
// u = &y
// fmt.Println(*u)
// } //  в таком случае ошибки не будет,  так как переменная на которую будет ссылаться u имеет значение.
//package main

//import "fmt"

//type Foo struct {
//bar string
//}

//func (f Foo) AddText(text string) string {
//return f.bar + " " + text //Структура Foo имеет метод AddText, в котором приемником является переменная f типа Foo.
//}

//func main() {

//f := Foo{
//bar: "Hi!",
//}

//fmt.Println(f.AddText("Text"))
//}

//Если какой-либо метод типа Type имеет получатель-указатель, то и все другие методы типа Type должны иметь должны иметь указатель в качестве получателя, даже те, для которых это не требуется.

package main

import "fmt"

type Home struct {
	IsOpenDoor bool
}

// OpenDoor - открывает дверь
func (h *Home) OpenDoor() {
	h.IsOpenDoor = true
}

func main() {

	home := Home{
		IsOpenDoor: false,
	}

	// По умолчанию дверь закрыта,
	// убедимся, что это так
	fmt.Println(home)

	// Вызывем метод OpenDoor().
	// Отрываем дверь
	home.OpenDoor()

	// Смотрим на результат.
	fmt.Println(home)
}

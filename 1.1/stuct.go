package main

import "fmt"

type Person struct {
	ID     int
	Name   string
	Addres string
}

type Account struct {
	ID      int
	Name    string
	Cleaner func(string) string
	Owner   Person
}

func main() {
	var acc Account = Account{
		ID:   1,
		Name: "rvasily",
	}
	fmt.Printf("%#v\n", acc)

	acc.Owner = Person{2, "Romanov Vasily", "Moscow"}
	fmt.Printf("%#v\n", acc)
}

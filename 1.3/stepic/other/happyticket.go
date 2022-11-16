package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	var a int
	a = num / 1000
	fmt.Println(a)
	var b int
	b = num % 1000
	//b = num % 10000 / 1000
	fmt.Println(b)
	var c int
	c = a / 100
	fmt.Println(c)
	var d int
	d = a / 10 % 10
	//d = num % 30000 / 3000
	fmt.Println(d)
	var i int
	i = a % 10 % 10
	fmt.Println(i)
	var j int
	j = b / 100
	//j = num % 30000 / 3000
	fmt.Println(j)
	s := b / 10 % 10
	fmt.Println(s)
	r := b % 10 % 10
	fmt.Println(r)
	onehalfticket := c + d + i
	fmt.Println(onehalfticket)
	twohalfticket := j + s + r
	fmt.Println(twohalfticket)
	if onehalfticket == twohalfticket {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

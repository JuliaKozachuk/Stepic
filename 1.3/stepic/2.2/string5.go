package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	c := strings.Split(s, "")
	fmt.Println(c)
	d := strings.Join(c, "*")
	fmt.Println(d)
}

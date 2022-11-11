package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b int

	fmt.Scan(&a, &b)

	c := math.Sqrt(float64(a*a + b*b))
	fmt.Println(c)
}

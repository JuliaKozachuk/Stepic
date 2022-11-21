package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var c int
	var sum int

	for i := 0; i <= N; i++ {
		fmt.Scan(&c)
		if c == 0 {
			sum++
		}
		fmt.Println(sum)

	}
}

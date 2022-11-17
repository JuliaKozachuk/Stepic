package main

import "fmt"

func main() {
	var a int

	for {
		fmt.Scan(&a)
		//var counter int
		//fmt.Scan(&counter)
		if a < 10 {
			continue
		} else if a > 100 {
			break
		} else {
			fmt.Println(a)
		}
	}
	//fmt.Println(a)

}

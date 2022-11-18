package main

import (
	"fmt"
	"strconv"
)

func main() {

	fn := func(i uint) uint {
		var number string
		//fmt.Println(fn)
		//return fn
		numberfn := strconv.FormatUint(uint64(i), 10)
		//func (i uint)uint{
		for _, a := range numberfn {
			if (a-'0')%2 == 0 && (a-'0') != 0 {
				number += string(a)

			}

		}
		j, _ := strconv.ParseUint(number, 10, 64)
		if j == 0 {
			return uint(100)
		}
		return uint(j)

	}
	fmt.Println(fn(727178))

}

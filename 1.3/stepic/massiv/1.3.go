package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	//for i, e := range array {
	//fmt.Println(array, i, e)
	//}
	for b := 0; b <= 0; b++ {
		l := array[0]
		c := array[1]
		d := array[2]
		e := array[3]
		j := array[4]
		//k:=array[5]

		if l > c && l > d && l > e && l > j {
			fmt.Println(l)

		}
		if c > l && c > d && c > e && c > j {
			fmt.Println(c)
		}
		if d > l && d > c && d > e && d > j {
			fmt.Println(d)
		}
		if e > l && e > c && e > d && e > j {
			fmt.Println(e)
		}
		if j > l && j > c && j > d && j > e {
			fmt.Println(j)
		}

	}

}

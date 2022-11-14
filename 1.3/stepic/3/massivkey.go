package main

import "fmt"

func main() {

	var arr [10]int

	// m:=make(map[int]int)
	// fmt.Scan(&m[1],&m[2],&m[3],&m[4],&m[5],&m[6],&m[7],&[8],&m[9],&m[10])
	// }
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	//fmt.Println(arr)

	m := make(map[int]int)
	for _, v := range arr {
		if value, ok := m[v]; ok {
			fmt.Println(value, " ")
		} else {
			m[v] = work(v)
			fmt.Println(m[v], " ")
		}
	}

}
func work(x int) int {
	return x + 1
}

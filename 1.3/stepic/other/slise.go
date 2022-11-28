package main

import "fmt"

func main() {
	var numone, numtwo int
	fmt.Scan(&numone, &numtwo)
	var i, j int
	for numone > 0 {
		numone = numone / 10
		i++
	}
	slice := make([]int, i, i)
	slicetwo := make([]int, j, j)
	for idx := range slice {
		if numone != 0 {
			i = i - 1
			idx = i
			slice[idx] = numone % 10
			numone = numone / 10
		}
	}
	for idxtwo := range slicetwo {
		if numtwo != 0 {
			j = j - 1
			idxtwo = j
			slicetwo[idxtwo] = numtwo % 10
			numtwo = numtwo / 10
		}

		for idx := range slice {
			if slice[idx] == numtwo {
				fmt.Println(slice[idx])
			} else {
				continue
			}
		}

	}
}

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	var del int
	fmt.Scan(&del)

	res, rank := 0, 1

	for num != 0 {
		resother := num % 10
		if resother != del {
			res += resother * rank
			rank *= 10
		}
		num /= 10
	}
	fmt.Println(res)

	// numstr := strconv.Itoa(num)
	// delstr := strconv.Itoa(del)

	// res := strings.Trim(numstr, delstr)
	// resin, _ := strconv.Atoi(res)

	// fmt.Println(resin)
}

// var N int
// fmt.Scan(&N)
// fmt.Printf("%b", N)
//}

//var res float64

// 	for {

// 		if math.Pow(2, res) <= float64(N) {
// 			fmt.Printf("%v ", math.Pow(2, res))
// 			res++
// 		} else {
// 			break
// 		}

// 	}

// }

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var c int
	fmt.Scan(&c)
	var d int
	fmt.Scan(&d)

	if n > 0 && n <= 10000 {
		for i := 1; i <= n; i++ {

			if i%c == 0 && i%d != 0 {
				fmt.Println(i)
				break

			}
		}

	}
}

// 		//} else if n%c != 0 && n%d == 0 {

// 		//}

// 	}

// }
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var n, c, d int
// 	fmt.Scan(&n, &c, &d)

// 	// secret condition
// 	if n > 0 && n < 10000 {
// 		for i := 1; i <= n; i++ {
// 			if i%c == 0 && i%d != 0 {
// 				fmt.Println("Условие выполнено")
// 				fmt.Println(i)
// 				break
// 			}
// 		}
// 	}
// }

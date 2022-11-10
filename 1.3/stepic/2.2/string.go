package main

import (
	"fmt"
	"strings"
)

func main() {
	var X string
	fmt.Scan(&X)
	var S string
	fmt.Scan(&S)
	a := strings.Contains(X, S)
	fmt.Println(a)
	i := strings.Index(X, S)
	fmt.Println(i)

}

// }
// type lines struct {
// 	X   string
// 	S   string
// 	Res int
// }

// func linesRes(l *lines) int {
// 	if a:=strings.Contains(l.X, l.S) {
// 		return a

// 	} else {
// 		return a-1
// 	}
//
// }

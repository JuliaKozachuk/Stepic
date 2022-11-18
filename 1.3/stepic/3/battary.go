package main

import (
	"fmt"
	"strings"
)

type battary struct {
	//empty string

	charged string
}

func (b *battary) String() string {
	empty := strings.Count(b.charged, "0")
	onedivision := strings.Count(b.charged, "1")
	totalcharge := "[" + strings.Repeat(" ", empty) + strings.Repeat("X", onedivision) + "]"
	return totalcharge

}

func main() {
	var a string
	fmt.Scan(&a)
	batteryForTest := &battary{charged: a}

	//empty:   a,
	fmt.Println(batteryForTest)

}

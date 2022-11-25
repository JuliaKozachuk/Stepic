package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)

	switch {
	case b == 1:
		fmt.Println(b, "korova")
	case b > 1 && b <= 4:
		fmt.Println(b, "korovy")
	case b >= 5 && b <= 20:
		fmt.Println(b, "korov")
	case b >= 21 && b%10 == 1:
		fmt.Println(b, "korova")
	case b >= 21 && (b%10 == 2 || b%10 == 3 || b%10 == 4):
		fmt.Println(b, "korovy")
	default:
		fmt.Println(b, "korov")
	}
}

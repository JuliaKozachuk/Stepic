package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var palalindrom string
	//fmt.Scan(&palalindrom)
	//if strings.Replace (palalindrom, palalindrom, palalindrom, -1){
	//fmt.Println("Палиндром")
	//}else{
	//fmt.Println("Нет")
	//}
	//pal := []string(palalindrom)
	palalindrom, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	pal := strings.TrimSpace(palalindrom)
	Runpal := []rune(palalindrom)
	length := len(Runpal)
	var backpal string

	//for i := range pal {
	//fmt.Print(i)

	//}
	//for a := 0; a < len(Runpal); a++ {
	//pal += string(pal[a])
	//if pal[a] != pal[a-1] {
	//fmt.Println("Палиндром")
	//} else {
	//fmt.Println("Нет")
	//}
	//
	//}
	//for a := range Runpal {
	//if Runpal[a] != pal[a] {
	//	fmt.Println("Палалиндром")
	//} else {
	//	fmt.Println("Нет")
	//}
	//}
	for a := length - 1; a >= 0; a-- {
		backpal += string(Runpal[a])
	}
	if pal != backpal {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}

}

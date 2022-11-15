package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//scanner,_ := bufio.NewScanner(os.Stdin).Re
	//text, _ = reader.ReadString('\n')
	s = strings.TrimSpace(s)
	//fmt.Println(s)
	number := strings.Split(s, ";")
	//fmt.Println(number)
	for i, j := range number {
		fmt.Println(i, j)
		j = strings.ReplaceAll(j, " ", "")
		//fmt.Println(j)
		j = strings.ReplaceAll(j, ",", ".")
		//fmt.Println(j)
		number[i] = j

	}
	number1, _ := strconv.ParseFloat(number[0], 64)
	number2, _ := strconv.ParseFloat(number[1], 64)
	res := number1 / number2
	fmt.Printf("%.4f\n", res)

}

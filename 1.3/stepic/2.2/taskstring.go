package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	//i, _ := strconv.Atoi(s)
	//fmt.Println(i)
	RuneI := []rune(s)
	max := string(RuneI[0])
	for i := 0; i < len(RuneI); i++ {
		if max < string(RuneI[i]) {
			max = string(RuneI[i])
		}
		//fmt.Println(max)
	}
	fmt.Println(max)
	//var max int
	//max != 1000
	//Irun := []rune(i)
	//for a := range Irun {
	//fmt.Println(a)
	//}
	//fmt.Println(max(i))

	//}

	//func max(i int)int{

	// var min int
	// min=0
	// }
	//i:=strconv.AppendInt(s)
	//a := strings.Split(s, "")
	//fmt.Println(a)
	//d := string.
	//fmt.Scan(d)

	//if a> 0 {
	//fmt.Println(strings.Split(s,string.Count a))

	//}
	//b :=
}

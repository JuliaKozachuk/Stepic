package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var b string
	//fmt.Scan(&s)
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//Rune:=strings.TrimSpace(s)
	RuneS := []rune(s)
	//fmt.Println(strings.Split(Rune))
	//b := strings.Replace(rs,rs [0,1,2,3,4,5,6,7,8,9,10],rs[1,3,5,7,9],-1) string
	//fmt.Println(a)
	for i := range RuneS {

		//a := strings.Split(i)
		fmt.Println(i)
		//fmt.Println(strings.Split(s,i))\
		if strings.Count(s, string(RuneS[i])) >= 1 {
			continue
		} else {
			s += string(RuneS[i])
		}
	}
	fmt.Println(string(RuneS[0:9]))

}

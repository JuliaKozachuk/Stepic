package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)
	textRune := []rune(text)
	if unicode.IsUpper(textRune[0]) && strings.HasSuffix(text, ".") {
		fmt.Printf("Right")
	} else {
		fmt.Println("Wrong")
	}

}

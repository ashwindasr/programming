package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("Enter a string: ")
	inputReader := bufio.NewReader(os.Stdin)
	s, _ := inputReader.ReadString('\n')
	// var s string
	// fmt.Scanln(&s)
	s = strings.ToLower(s)
	pos := strings.IndexRune(s, 'a')
	if s[0] == 'i' && s[len(s)-3] == 'n' && pos > 0 && pos < (len(s)-2) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}

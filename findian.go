package main

import s "strings"

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var inputText string
	fmt.Printf("enter a string ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputText = scanner.Text()
	}

	inputText = s.ToLower(inputText)
	if s.HasPrefix(inputText, "i") == true && s.HasSuffix(inputText, "n") == true && s.Contains(inputText, "a") == true {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}
}

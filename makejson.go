// Write a program which prompts the user to first enter a name, and then enter an address.
// Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	name := getInput("Enter a name")
	address := getInput("Enter an address")

	person := make(map[string]string) //creating a map
	person["name"] = name
	person["address"] = address

	jsonObj, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonObj))
}

func getInput(prompt string) string { //input is the prompt text, output is the value of the prompt text
	var input string

	fmt.Printf(prompt + ": ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text() //scanning the text on command line to inout variable
	}

	return input
}

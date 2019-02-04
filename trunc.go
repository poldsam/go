package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Enter a floating number:")
	var floatingNum float64
	fmt.Scanln(&floatingNum)
	fmt.Println(math.Round(floatingNum))

}

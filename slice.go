package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s1 := make([]int, 3) //make a slice
	var counter int

	for {
		var k string //capture command line text
		fmt.Printf("Enter an Integer: ")
		fmt.Scan(&k)

		intK, _ := strconv.Atoi(k) //Atoi to convert string k to int)

		if k != "X" && k != "x" {
			if counter < 3 {
				s1[counter] = intK
			} else {
				s1 = append(s1, intK)
			}

			s2 := make([]int, len(s1)) //create empty slice with length of s1
			copy(s2, s1)               //copy s1(source) values to s2(destination)
			sort.Ints(s2)              //sort s2 values
			fmt.Println(s2)
			counter++
		} else {
			break
		}
	} //end of for loop
} //end of function

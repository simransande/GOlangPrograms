package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var headcount int = 0
	var tailcount int = 0
	var per int = 0

	var perr int = 0

	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanf("%d", &input)
	var a [20]int

	for i := 1; i <= input; i++ {
		a[i] = rand.Intn(input)
		if a[i] < input/2 {
			headcount++
		} else {
			tailcount++
		}

	}

	per = (headcount / input) * 100
	perr = (tailcount / input) * 100

	fmt.Print(per)
	fmt.Print(perr)
	// output := input * 2

	// fmt.Println(output)
}

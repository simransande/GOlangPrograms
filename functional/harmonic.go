package main

import "fmt"

func main() {

	var n int
	var sum float32 = 0
	fmt.Print("Enter a positive integer new  : ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ { // assigning 1 to i
		var m = float32(i)
		sum += (1 / m) // sum = sum + i

		if i != n {
			fmt.Printf("1/ %v +", i)
			// fmt.Print("sadfsdg ")
			// fmt.Printf("intSlice \tLen: %v \tCap: %v\n", i, i)

			// fmt.Printf("1123  %v \n ", i)

		} else {
			fmt.Printf("1/%v\n", i)
		}

	}
	fmt.Printf("harmonic number of %v is %v\n", n, sum)

}

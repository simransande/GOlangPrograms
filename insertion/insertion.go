package main

import (
	"fmt"
)

func main() {

	sorted := array(20)

	insertionsort(sorted)
	fmt.Print("\n--- Sorted ---\n\n", sorted, "\n")
}
func array(size int) []int {

	fmt.Println(`Enter the number of integers`)
	var n int
	if m, err := Scan(&n); m != 1 {
		panic(err)
	}
	fmt.Println(`Enter the integers`)
	all := make([]int, n)
	ReadN(all, 0, n)
	fmt.Println(all)
	return all
}

func ReadN(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := Scan(&all[i]); m != 1 {
		panic(err)
	}
	ReadN(all, i+1, n-1)
}

func Scan(a *int) (int, error) {
	return fmt.Scan(a)
}

func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}

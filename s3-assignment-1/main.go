package main

import "fmt"

func main() {
	xs := generateSlice()

	for _, v := range xs {
		if v%2 == 0 {
			fmt.Printf("%d is even\n", v)
		} else if v%2 == 1 {
			fmt.Printf("%d is odd\n", v)
		}
	}
}

func generateSlice() []int {
	xs := []int{}

	for i := 0; i < 11; i++ {
		xs = append(xs, i)
	}
	return xs
}

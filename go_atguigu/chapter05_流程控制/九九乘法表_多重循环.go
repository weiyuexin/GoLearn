package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d   \t", i, j, i*j)
		}
		fmt.Println()
	}
}

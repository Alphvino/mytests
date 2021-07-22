package main

import "fmt"

func main() {
	var n uint8
	fmt.Scanln(&n)
	for i := uint8(1); i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}

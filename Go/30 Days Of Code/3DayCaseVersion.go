package main

import "fmt"

func main() {
	var n uint8
	fmt.Scanln(&n)
	even := n%2 == 0
	switch {
	case even && n >= 2 && n <= 5:
		fmt.Println("Not Weird")
	case even && n >= 6 && n <= 20:
		fmt.Println("Weird")
	case even && n > 20:
		fmt.Println("Not Weird")
	default:
		fmt.Println("Weird")
	}
}

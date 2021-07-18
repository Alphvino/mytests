package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	if scanner == scanner {
		var g uint64
		var f float64
		var k string
		var c string

		_, err := fmt.Scan(&g, &f, &c)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}

		k, err = bufio.NewReader(os.Stdin).ReadString('\n')

		fmt.Println(i + g)
		fmt.Printf("%.1f\n", d+f)
		fmt.Println(s + c + " " + k)
	}

}

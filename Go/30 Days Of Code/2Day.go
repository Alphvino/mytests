package main
import (
	"fmt"
	"math"
)

func main() {
	var mealprice, tipercent, taxpercent float64
	fmt.Scan(&mealprice, &tipercent, &taxpercent)
	tip := (tipercent / 100) * mealprice
	tax := (taxpercent / 100) * mealprice
	total := mealprice + tip + tax
	fmt.Println(math.Round(total))
}

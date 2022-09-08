package main

import (
	"fmt"
	"math"
)

func main() {
	var ans string

	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range values {
		if math.Mod(float64(value), float64(2)) == 0 {
			ans = "Even"
		} else {
			ans = "Odd"
		}

		fmt.Printf("%v is %v\n", value, ans)
	}
}

package main

import (
	"fmt"
)

func main() {
	var numbers, a, sum int
	fmt.Scan(&numbers)

	for i := 0; i != numbers; i++ {
		fmt.Scan(&a)
		if a >= 10 && a < 100 && a%8 == 0 {
			sum = sum + a
		}

	}
	fmt.Println(sum)
}

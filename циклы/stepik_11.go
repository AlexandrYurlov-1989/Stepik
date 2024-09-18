package main

import (
	"fmt"
)

func main() {

	var x, p, y int

	fmt.Scan(&x, &p, &y)
	i := 0

	for x <= y {
		x = (x*p)/100 + x

		i++
	}
	fmt.Println(i)
}

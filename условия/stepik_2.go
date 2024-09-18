package main

import "fmt"

func main() {
	var number, hundreds, tens, units int
	fmt.Scan(&number)

	hundreds = number / 100
	tens = (number / 10) % 10
	units = number % 10

	switch {

	case hundreds == tens || hundreds == units || tens == units:
		fmt.Println("NO")
	default:
		fmt.Println("YES")
	}
}

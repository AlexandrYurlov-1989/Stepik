package main8

import (
	"fmt"
)

func main8() {
	var number, a, max_limit int
	number = 1
	for number != 0 {
		number = 0
		fmt.Scan(&number)
		if number > a {
			a = number
			max_limit = 1

		} else if number == a && number != 0 {
			max_limit++

		}

	}
	fmt.Println("максимальный ", max_limit)
}

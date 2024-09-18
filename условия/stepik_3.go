package main

import "fmt"

func main() {
	var number, hundreds int
	fmt.Scan(&number) //1234

	switch {
	case number <= 9999:
		hundreds = number / 1000
	case number <= 999:
		hundreds = number / 100
	case number <= 99:
		hundreds = number / 10
	case number <= 9:
		hundreds = number / 1
	}
	fmt.Println(hundreds)

}

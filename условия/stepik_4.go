package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number) //1234

	if number <= 10000 && number > 0 {
		switch {

		case number%4 == 0 && number%100 != 0:
			fmt.Print("YES")

		case number%400 == 0:
			fmt.Print("YES")

		default:
			fmt.Print("NO")
		}
	}
}

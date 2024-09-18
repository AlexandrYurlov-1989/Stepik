package main

import "fmt"

func main() {
	num := 0
	if num > 0 {
		fmt.Println("Число больше нуля")
	} else if num < 0 {
		fmt.Println("Число меньше нуля")
	} else if num == 0 {
		fmt.Println("Число равно 0")
	}

}

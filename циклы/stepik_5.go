package main

import "fmt"

func main() {
	a := 0
	for i := 1; i < 11; i++ {
		a = i * i
		fmt.Println(a)
	}

}

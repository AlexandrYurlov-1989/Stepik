package main

import (
	"fmt"
)

func main() {
	var n int

	for {
		fmt.Scanln(&n)

		if n < 10 {
			continue // Пропускаем числа меньше 10
		}
		if n > 100 {
			break // Прекращаем считывание, если число больше 100
		}

		fmt.Println(n) // Выводим число, если оно в диапазоне [10, 100]
	}
}

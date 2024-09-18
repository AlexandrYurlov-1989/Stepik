package main // уже готово, проверять после изучения курса о массивах.

import (
	"fmt"
)

func main() {
	var num1, num2 string
	fmt.Scan(&num1, &num2)

	// Создаем множество для хранения цифр второго числа
	digitsMap := make(map[rune]bool)

	// Заполняем множество цифрами из второго числа
	for _, digit := range num2 {
		digitsMap[digit] = true
	}

	// Срез для хранения общих цифр
	var commonDigits []rune

	// Проверяем цифры первого числа
	for _, digit := range num1 {
		if digitsMap[digit] {
			commonDigits = append(commonDigits, digit)
		}
	}

	// Выводим результат
	for i, digit := range commonDigits {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(string(digit))
	}
	fmt.Println()
}

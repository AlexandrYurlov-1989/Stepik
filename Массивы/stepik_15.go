package main

import (
	"fmt"
)

func main() {
	// Создаем массив из 10 элементов
	var workArra [10]uint8

	// Читаем 10 целых положительных чисел
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArra[i])
	}
	// Читаем 3 пары индексов и меняем местами элементы массива
	fmt.Println(workArra)

	for i := 0; i < 3; i++ {
		var index1, index2 uint8
		fmt.Scan(&index1, &index2)
		workArra[index1], workArra[index2] = workArra[index2], workArra[index1]
	}
	// Выводим элементы массива через пробел

	for _, value := range workArra {
		fmt.Print(value, " ")
	}
}

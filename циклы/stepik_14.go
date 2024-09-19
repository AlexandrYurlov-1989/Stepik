	var workArray [10]uint8

	// Читаем 10 целых положительных чисел
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}

	// Читаем 3 пары индексов и меняем местами элементы массива
	for i := 0; i < 3; i++ {
		var index1, index2 int8
		fmt.Scan(&index1, &index2)
		workArray[index1], workArray[index2] = workArray[index2], workArray[index1]
	}

	// Выводим элементы массива через пробел
	for _, value := range workArray {
		fmt.Print(value, " ")
	}

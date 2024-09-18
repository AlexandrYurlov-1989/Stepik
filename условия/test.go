package main

import "fmt"

func main() {
	var name string
	var age uint8
	fmt.Println("Как тебя завут")
	fmt.Scan(&name)
	fmt.Println("Привет " + name + "!")
	fmt.Println("сколько тебе лет?")
	fmt.Scan(&age)
	fmt.Println("Тебе " + fmt.Sprint(age) + " лет!!!")
}

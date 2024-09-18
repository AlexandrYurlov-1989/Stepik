package main

import (
 "fmt"
)

func main() {
 var number int
 fmt.Scan(&number) //613244

 // Определяем количество цифр в первой части
 firstPart := number / 1000  // Получаем 613
 secondPart := number % 1000 // Получаем 244
 fmt.Println(firstPart)
 fmt.Println(secondPart)

 firstPart1 := firstPart / 100
 fmt.Println(firstPart1)
 firstPart2 := (firstPart / 10) % 10
 fmt.Println(firstPart2)
 firstPart3 := firstPart % 10
 fmt.Println(firstPart3)
 sum1 := firstPart1 + firstPart2 + firstPart3
 fmt.Println(sum1)

 secondPart1 := secondPart / 100
 fmt.Println(secondPart1)
 secondPart2 := (secondPart / 10) % 10
 fmt.Println(secondPart2)
 secondPart3 := secondPart % 10
 fmt.Println(secondPart3)
 sum2 := secondPart1 + secondPart2 + secondPart3
 fmt.Println(sum2)

 if sum1 == sum2 {
  fmt.Println("YES")
 } else {
  fmt.Println("NO")
 }

}

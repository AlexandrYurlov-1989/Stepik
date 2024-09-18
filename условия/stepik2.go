package main2

import (
 "fmt"
)

func main2() {
 var number, hundreds int
 fmt.Scan(&number) //1234

 switch {

 case number < 0:
  break
 case number <= 9:
  hundreds = number / 1
  fmt.Println(hundreds)
 case number <= 99:
  hundreds = number / 10
  fmt.Println(hundreds)
 case number <= 999:
  hundreds = number / 100
  fmt.Println(hundreds)
 case number <= 9999:
  hundreds = number / 1000
  fmt.Println(hundreds)
 case number == 10000:
  hundreds = number / 10000
  fmt.Println(hundreds)
 default:
  break
 }

}

package main1

import "fmt"

func main1() {
 var a int
 fmt.Scan(&a)

 switch {

 case a == 0:
  fmt.Print("Ноль")

 case a > 0:
  fmt.Print("Число положительное")

 case a < 0:
  fmt.Print("Число отрицательное")
 }

}

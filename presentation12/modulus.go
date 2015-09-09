package main

import "fmt"

func main() {

  var input1 int
  fmt.Print("Enter a number from 1 - 10")
  fmt.Scan(&input1)
  
  var input2 int
  fmt.Print("Enter a  second number from 1 - 10")
  fmt.Scan(&input2)
  
  var result int
  result = input1 % input2
  fmt.Println("Dividing the first number ", input1, " by the second number ", input2, " leaves a remainder of ", result)
  
  
  
}
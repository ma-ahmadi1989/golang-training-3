package main

import "fmt"

func main() {

	var num1 int16
	var num2 int16

	fmt.Print("Enter ur number: ")
	fmt.Scan(&num1, &num2)

	result := num1 + num2

	if result > 0 && result/10 == 0 {
		println("True")
	}else{
		println("False")
	}
}
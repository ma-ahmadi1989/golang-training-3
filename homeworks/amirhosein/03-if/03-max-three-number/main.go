package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Print("Enter the first number:")
	fmt.Scan(&num1)

	fmt.Print("Enter the two number:")
	fmt.Scan(&num2)

	fmt.Print("Enter the three number:")
	fmt.Scan(&num3)

	if num1 == num2 || num1 == num3 || num3 == num2 {
		fmt.Println("The entered numbers are repeated")
	} else if num1 > num2 && num1 > num3 {
		fmt.Println("the first number is greater:", num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Println("the two number is greater:", num2)
	} else {
		fmt.Println("the three number is greater:", num3)
	}

}

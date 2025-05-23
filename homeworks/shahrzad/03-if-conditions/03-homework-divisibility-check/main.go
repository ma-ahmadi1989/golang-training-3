package main

import(

	"fmt"
)

func main(){

	var firstNum, secondNum int

	fmt.Print("Enter your first number: ")
	fmt.Scan(&firstNum)

	fmt.Print("Enter your second number: ")
	fmt.Scan(&secondNum)

	if firstNum == 0 || secondNum == 0 {
		fmt.Println("Error: Division by zero is not allowed.")
		return
	}

	if firstNum > secondNum{
		if firstNum % secondNum == 0{
			fmt.Printf("Yes, %d is a divisor of %d.", secondNum, firstNum)
		}else{
			fmt.Printf("No, %d is not a divisor of %d.", secondNum, firstNum)
		}

	}else{
		if secondNum % firstNum == 0{
			fmt.Printf("Yes, %d is a divisor of %d.", firstNum, secondNum)
		}else{
			fmt.Printf("No, %d is not a divisor of %d.", firstNum, secondNum)
		}
	}
}
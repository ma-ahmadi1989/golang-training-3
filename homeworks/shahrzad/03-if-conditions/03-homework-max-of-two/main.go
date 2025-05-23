package main

import(
	"fmt"
)

func main(){

	var number1, number2 float32
	
	fmt.Print("Enter your first number: ")
	fmt.Scan(&number1)
	
	fmt.Print("Enter your second number: ")
	fmt.Scan(&number2)

	if number1 > number2{
		fmt.Printf("%0.2f is the greater number.", number1)
	}else{
		fmt.Printf("%0.2f is the greater number.", number2)
	}

}
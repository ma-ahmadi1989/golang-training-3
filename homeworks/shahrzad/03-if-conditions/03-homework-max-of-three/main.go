package main 

import(

	"fmt"
)

func main(){

	var numberA, numberB, numberC float32

	fmt.Print("Enter your first number: ")
	fmt.Scan(&numberA)

	fmt.Print("Enter your second number: ")
	fmt.Scan(&numberB)

	fmt.Print("Enter your third number: ")
	fmt.Scan(&numberC)

	if numberA > numberB && numberA > numberC{
		fmt.Printf("%0.2f is the largest number.", numberA)
	}else if numberB > numberA && numberB > numberC{
		fmt.Printf("%0.2f is the largest number.", numberB)
	}else{
		fmt.Printf("%0.2f is the largest number.", numberC)
	}
}
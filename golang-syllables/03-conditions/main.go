package main

import "fmt"

func main() {

	a := 2

	remaining := a % 2

	// Operands:
	// == equal
	// != not equal
	// > greater than
	// >= equal or greater than
	// < less than
	// <= equal or less than
	// && AND
	// || OR

	if remaining == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	if remaining != 0 {
		fmt.Println("The number is odd")
	}

	// ----------------------
	b := 100

	if b == 100 {
		fmt.Println("it is literally 100!")
	} else if b > 100 {
		fmt.Println("it is greater than 100!")
	} else {
		fmt.Println("it is less than 100!")
	}

	if b == 100 {
		fmt.Println("it is literally 100!")
	} else if b > 100 {
		fmt.Println("it is greater than 100!")
	} else if b < 100 {
		fmt.Println("it is less than 100!")
	}

	// ------------
	// description: if number is even and greater than 35

	// answer 1
	a1 := 16
	a1Remaining := a1 % 2

	if a1Remaining == 0 {
		if a1 > 35 {
			fmt.Println("the number is even and greater than 35")
		} else {
			fmt.Println("the number is even but less than 35")
		}
	} else {
		fmt.Println("The number is not even")
	}

	// answer 2
	a2 := 16
	a2Remaining := a2 % 2

	if a2Remaining == 0 || a2 > 35 {
		fmt.Println("the number is even and greater than 35")
	} else {
		fmt.Println("the number is not what we want!")
	}

	// -------------------
	// if is between 1-10 and is even or is between 30-40 and is odd
	m := 15

	mRemaining := m % 2

	if (m > 1 && m < 10 && mRemaining == 0) || (m > 30 && m < 40 && mRemaining != 0) {
		fmt.Println("the number is even and greater than 35")
	} else {
		fmt.Println("the number is not what we want!")
	}

}

package main

import "fmt"

func main() {
	var m int
	fmt.Print("enter the input: ")
	fmt.Scan(&m)
	mRemaining := m % 2
	m5Remaining := m % 5
	if m > 0 {
		fmt.Print("it is greater than 0, ")
		if m5Remaining == 0 {
			fmt.Print("x % 5 = 0, ")
			if mRemaining == 0 {
				fmt.Print("it is even")
			}
		} else if mRemaining == 0 {
			fmt.Print("it is even")
		}
	} else if m == 0 {
		fmt.Println("it is zero")
	} else {
		fmt.Println("it is less than 0")
	}
}

package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("enter number one :")
	fmt.Scan(&a)
	fmt.Print("enter number two :")
	fmt.Scan(&b)
	if a < b {
		if b%a == 0 {
			fmt.Printf("%d is a divisor %d", a, b)
		}
	} else if a > b {
		if a%b == 0 {
			fmt.Printf("%d is a divisor %d", b, a)
		}
	} else if a == b {
		fmt.Printf("%d is a divisor %d", b, a)
	} else {
		fmt.Println("non of them is divisor")
	}

}

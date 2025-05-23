package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("enter the number : ")
	fmt.Scan(&a)
	fmt.Print("enter the number as power : ")
	fmt.Scan(&b)
	var result int
	result = 1
	for i := 0; i < b; i++ {
		result = result * a
	}
	fmt.Printf("result : %d", result)
}

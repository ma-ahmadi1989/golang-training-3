package main

import "fmt"

func main() {
	var lenght float64
	fmt.Print("enter lenght :")
	fmt.Scan(&lenght)

	var width float64
	fmt.Print("enter width :")
	fmt.Scan(&width)

	area := width * lenght

	fmt.Printf("rectangle area : %.2f", area)
}

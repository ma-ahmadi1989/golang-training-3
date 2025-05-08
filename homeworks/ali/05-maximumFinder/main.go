package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Print("enter numer one :")
	fmt.Scan(&a)
	fmt.Print("enter numer two :")
	fmt.Scan(&b)
	fmt.Print("enter numer three :")
	fmt.Scan(&c)
	max := a
	if a < b && c < b {
		max = b
	} else if a < c {
		max = c
	}
	fmt.Printf("%0.2f is the greatest among all", max)
}

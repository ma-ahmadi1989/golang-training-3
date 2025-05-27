package main

import "fmt"

func main() {
	var width, length int
	fmt.Print("Enter width:")
	fmt.Scanln(&width)

	fmt.Print("Enter length:")
	fmt.Scanln(&length)

	perimete := 2 * (width + length)
	fmt.Printf("Rectangle perimete:%d \n", perimete)

	area := width * length
	fmt.Printf("Rectangle area:%d", area)
}

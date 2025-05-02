package main

import "fmt"

func main() {

	fmt.Println("Hello world!")

	// How to define a variabl

	// type 1
	// var name string
	// name = "Mahmoud"

	// // type 2
	// familyName := "Ahmadi"

	// // type 3
	// var fatherName string = "Davoud"

	// // type 4
	// name, familyName, fatherName := "Mahmoud", "Ahmadi", "Davoud"

	// // type 5
	// var (
	// 	name       string
	// 	familyName string
	// 	fatherName string
	// )
	// name = "Mahmoud"
	// familyName = "Ahmadi"
	// fatherName = "Davoud"

	// // type 6
	var (
		name       string = "Mahmoud"
		familyName string = "Ahmadi"
	)

	age := 37
	var Score float32 = 8 // the max is 10
	var (
		isMale bool
	)
	isMale = true

	var (
		isMarried bool = true
	)

	fmt.Println("name: ", name, " Family Name: ", familyName)

	fmt.Printf("name: %s, FamilyName: %s, Age: %d, Score: %.2f, IsMale: %t, Is Married: %t \n",
		name,
		familyName,
		age,
		Score,
		isMale,
		isMarried)

	age = 2 + age
	age += 2

	fullName := name + " " + familyName

	var brotherName string
	fmt.Printf("Enter your brother name: ")
	fmt.Scanf("%s", &brotherName)

	fmt.Printf("Your brother name is: %s \n", brotherName)
	fmt.Printf("Your full name is: %s \n", fullName)

}

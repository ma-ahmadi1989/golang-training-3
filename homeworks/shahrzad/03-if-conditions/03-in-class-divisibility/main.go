package main

import "fmt"

func main() {

	var num int16

	fmt.Print("Enter ur number: ")
	fmt.Scan(&num)


	if (num>0) && (num%10 == 0) {
		println("True")
	}else{
		println("False")
	}
}
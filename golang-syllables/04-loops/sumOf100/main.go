package main

import "fmt"

func main() {

	sum := 0

	for j := 0; j < 2; j++ {
		for i := 1; i <= 100; i++ {
			sum += i
			if sum < 1000 {
				break
			}
		}
	}

	fmt.Println(sum)

}

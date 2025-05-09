package main

import (
	"fmt"
	"time"
)

func main() {

	// step := 0

	// for step < 100 {
	// 	fmt.Println("step: ", step)
	// 	step++
	// }

	// for i := 0; i < 200; i++ {
	// 	fmt.Println("i: ", i)
	// }

	// num := 0

	// for {
	// 	if num == 100 {
	// 		break
	// 	}
	// 	time.Sleep(300 * time.Millisecond)
	// 	fmt.Println(num)
	// 	num++
	// }

	// fmt.Println("End of loop")

	j := 0

	for {
		if (j % 2) == 0 {
			j++
			continue
		}
		fmt.Println(j)
		j++
		time.Sleep(300 * time.Millisecond)
	}
}

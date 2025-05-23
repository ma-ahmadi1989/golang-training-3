package main

import "fmt"

func main() {
    var number int

    fmt.Print("Enter a number: ") 
    fmt.Scan(&number)

    fmt.Printf("Multiplication table of %d", number) 
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d * %d = %d\n", number, i, number*i)
    }
}

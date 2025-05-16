func main() {
	var n int
	fmt.Print("how many number ? : ")
	fmt.Scan(&n)

	sum := 0
	var num int

	for i := 0; i < n; i++ {
		fmt.Printf("enter number %d: ", i+1)
		fmt.Scan(&num)
		sum += num
	}

	average := float64(sum) / float64(n)
	fmt.Println("average:", average)
}

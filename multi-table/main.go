package main

import "fmt"

func main() {

	var number int
	var limit int

	fmt.Println("Give me number to multiply")
	fmt.Scanln(&number)

	fmt.Println("What is the limit?")
	fmt.Scanln(&limit)

	multi := 1

	for multi <= limit {
		output := number * multi
		fmt.Printf("%d x %d = %d\n", number, multi, output)
		multi = multi + 1
	}
}

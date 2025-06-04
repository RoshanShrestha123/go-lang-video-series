package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	var age int

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello, What is your name?")
	fmt.Scanln(&name)

	fmt.Println("What is your age?")
	fmt.Scanln(&age)

	fmt.Println("So where do you live?")
	address, _ := reader.ReadString('\n')

	fmt.Printf("Oh, so your name is %s and your age is %d. Really? you look older than that!\n", name, age)
	fmt.Printf("You live in %s", address)
}

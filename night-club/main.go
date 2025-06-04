package main

import "fmt"

func main() {
	var age int
	fmt.Println("Welcome to the LOD club")
	fmt.Println("Bouncer 1: Can i see your id? And what is your age?")

	fmt.Scanln(&age)

	if age == 18 {
		fmt.Println("Ok you are allow to go in")
	} else {
		fmt.Println("You are just kid man, go to school!!")
	}

}

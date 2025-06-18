package main

import (
	"math/rand"
)

func main() {
	num1, num2, user_name := randomNumber()

	println(num1, num2, user_name)
}

func randomNumber() (int, int, string) {
	num1 := rand.Intn(10)
	num2 := rand.Intn(10)

	name := "roshan"

	return num1, num2, name
}

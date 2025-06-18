package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(10) + 1
	var un int

	var health = 3

	for health > 0 {
		fmt.Println("👀 Hey man, guess the number i am thinking from 1-10")
		fmt.Scanln(&un)

		if un == n {
			fmt.Println("✅ Yes!, you guessed it right")
			return
		} else if un > n {
			fmt.Println("❌ Guess lower")
		} else {
			fmt.Println("❌ Guess higher")
		}

		health = health - 1

	}

}

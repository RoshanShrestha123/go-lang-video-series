package main

import "fmt"

func main() {

	day := "wednesday"

	switch day {
	case "monday":
		fmt.Println("Oh aja monday? kaam suru garam")

	case "tuesday":
		fmt.Println("Oh aja tuesday? alxi layo")

	case "wednesday":
		fmt.Println("Oh aja wednesday? alxi layo")
		fallthrough

	case "thursday":
		fmt.Println("Oh aja thursday? oe kati xito thursday ayo lol")

	case "friday":
		fmt.Println("half xutti")

	case "saturday", "sunday":
		fmt.Println("bida")

	default:
		fmt.Println("Yesto ni hunxa week???")

	}
}

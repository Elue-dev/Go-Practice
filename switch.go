package main

import "fmt"

func main() {
	animal := "Dog"
	var action string

	switch animal {
	case "Cat":
		action = "Meow"
	case "Dog":
		action = "Bark"
	case "Horse":
		action = "Neigh"
	default:
		action = "Nothing"
	}

	fmt.Println(action)
}
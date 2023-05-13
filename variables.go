package main

import "fmt"

func main() {
	var favBook string
	var condition bool
	var myNumber int

	favBook = "Beloved"
	condition = true
	myNumber = 10

	fmt.Println(favBook, condition, myNumber)

	// Compound Creation
	var favNumber, favChocolate, favTeam = 23, "Snickers", "Chelsea"
	fmt.Println(favNumber, favChocolate, favTeam)

	// Block Creation
	var (
		otherfavNumber = 10
		otherfavChocolate = "Nutella"
		otherfavTeam = "Warriors"
	)

	fmt.Println(otherfavNumber, otherfavChocolate, otherfavTeam)

	favPlayer := "Stephen Curry"
	fmt.Println(favPlayer)

	pet1, pet2, pet3 := "Dog", "Cat", "Parrot"
	fmt.Println(pet1, pet2, pet3)

	//Constants
	const myName = "Wisdom"

	// myName = "Dubem" - won't work
}
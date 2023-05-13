package main

import "fmt"

func main() {
	name, age := "Elue", 25
	var isInvited = name == "Elue" && age > 23
	var decision string

	if isInvited {
		decision = "You are invited to the"
	} else {
		decision = "You are not invited"
	}

	fmt.Println(decision)
}
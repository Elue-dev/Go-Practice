package utimainls

import "fmt"

func main() {
	// Arithmetic Opeartors
	const num1 = 4
	const num2 = 3
	const num3 = 4

	const sum = num1 + num2
	const difference = num1 - num2
	const quotent = num1 / num2
	const product = num1 * num2
	const remainder = num2 % 2

	fmt.Println(sum, difference, quotent, product, remainder)

	// Relational Operators
	const result  = num1 > num2
	const result2  = num2 > num1
	const result3 = num1 == num3
	const result4 = num1 != num2
	fmt.Println(result, result2, result3, result4)

	// Logical Operators
	name, age := "Elue", 25
	var isInvited = name == "Elue" && age > 23
	var isInvited2 = !( name == "Elue") && (age >=21 && age < 40)
	
	fmt.Println(isInvited, isInvited2)

}
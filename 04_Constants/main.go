package main

import (
	"fmt"
	"math"
)

func main() {
	//	A constant in golang is a variable whose value can't be modified once it's declared

	const num = 12
	//num = 13 --> WRONG
	fmt.Println(num)

	//	Declaring a group of constants

	const (
		name    = "Nitish"
		age     = 20
		country = "India"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	//	The value of a constant should be known at compile time.
	//	Hence, it cannot be assigned to a value returned by a function call since the function call takes place at run time.

	var num2 = math.Sqrt(25)
	//const num3 = math.Sqrt(36) --> WRONG

	fmt.Println(num2)

	const n = "Sam"
	var nameee = n
	fmt.Printf("type %T value %v", nameee, nameee)
}

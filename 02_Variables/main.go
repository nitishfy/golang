package main

import "fmt"

func main() {
	//	A variable is just a name give to the memory location where the data is stored

	//	Variable Declaration methods

	//	Variable declaration by specifying type

	var age int = 25
	fmt.Println("My age is", age)

	//	Variable declaration without specifying type

	var name = "Nitish"
	fmt.Println("My name is", name)

	//	Variable declaration without using 'var' keyword --> := will be used --> Short hand declaration

	marks := 12
	fmt.Println("My marks are", marks)

	//	What if no initial value is assigned to a variable? --> default value will be considered then

	var rollNum int
	fmt.Println(rollNum) // 0

	// Multiple variable declaration

	length, breadth := 12, 12
	//var length, breadth = 12,12
	fmt.Println(length, breadth)

	//			or

	var (
		isPresent   = false
		lovesCoding = true
		followers   = 1500
	)

	fmt.Println("Is Nitish present?", isPresent)
	fmt.Println("Do Nitish loves coding?", lovesCoding)
	fmt.Printf("Nitish has %v followers on Twitter LOL", followers) // using printf here

	//	Printf don't add a new line whereas Println adds a new line

}

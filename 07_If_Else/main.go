package main

import "fmt"

func main() {
	fmt.Println("Enter the number")
	var num1 int
	fmt.Scanln(&num1) // taking user input

	if num1%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// another way

	if num := 16; num%2 == 0 { // assignment-statement;condition
		fmt.Println("Even")
	}

	// another way (Preferred)
	var num2 int
	fmt.Scanln(&num2)
	if num2%2 == 0 {
		fmt.Println("Even")
		return
	}
	fmt.Println("Odd")
}

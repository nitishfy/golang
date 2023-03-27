package main

import "fmt"

func main() {
	a, b := 12, 12
	result1 := addition(a, b)
	result2 := subtraction(a, b)
	fmt.Printf("Addition of %v and %v gives %v and difference gives %v", a, b, result1, result2)
	fmt.Println() // adds new line

	// calling Rectangle function
	area, peri := Rectangle(12, 12)
	fmt.Println(area, peri)

	//	calling NamedValue function
	fmt.Println(NamedValue(12, 12))

	// using blank identifier(_) to discard a value
	area1, _ := Rectangle(12, 12)
	fmt.Println(area1) // only area1 will be printed, peri has been discarded using _

}

// addition Function to add two numbers
func addition(a, b int) int { // a,b are parameters of type int, other int is the return type
	sum := a + b
	return sum
}

// subtraction Function to subtract two numbers
func subtraction(a, b int) int {
	difference := a - b
	return difference
}

// Constructing function of multiple return types

// Rectangle Function to calculate area and perimeter of rectangle
func Rectangle(length, breadth int) (int, int) {
	area := length * breadth
	perimeter := (length + breadth) * 2
	return area, perimeter
}

// Return named values from a function

// NamedValue returns names
func NamedValue(num1, num2 int) (multiplication, division int) {
	multiplication = num1 * num2
	division = num1 / num2
	return // no explicit return value to be provided in NamedValue functions
}

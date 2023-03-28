package learnpackage

import "fmt"

func Addition(a, b int) int {
	sum := a + b
	return sum
}

// init is a special function in Go that has no return type, paramter and called automatically at the beginning of a program
func init() {
	fmt.Println("learnpackage init function called")
}

// Here's how the program structures for an init package

/*
1. Imported packages are initialised first. Hence, their init method is called first
2. Package level variables are initialised next
3. init function of main package is called next
4. main function is called at last
*/

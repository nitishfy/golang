package main

import (
	"fmt"
	"learnpackage/demo"
	"learnpackage/learnpackage"
)

func init() {
	fmt.Println("main package init function called")
}

var a, b = 12, 12

func main() {
	fmt.Println("Importing the Addition function from learnpackage directory")
	result2 := demo.Subtraction(a, b)
	result := learnpackage.Addition(a, b)
	fmt.Println(result)
	fmt.Println(result2)

}

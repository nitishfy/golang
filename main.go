package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

/*
Important points to note here:

1. There are 4 ways to run a go program --> go build, go run, go install & go playground.

2. Every go file should start with 'package name' statement.

3. Understanding each line of the code above:

 - package main is the package that has been used here. Packages are used to provide code reusability
 - import statement is used to import a package. Here, fmt package is imported to print the text to standard output.
 - func main is a special function from where execution of a program starts.
 - package.Function() is the syntax to call a function in a package. For eg. Println is the function in fmt package.
*/

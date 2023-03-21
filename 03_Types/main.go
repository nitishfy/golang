package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var isTrue bool = true // Type --> bool(boolean) --> Either True or False
	fmt.Println(isTrue)

	var number int = 12 // Type --> int(Integer) --> there are various types of int exist like int8,int16,int32,int64
	fmt.Println(number)

	//	The difference between all these int types is their size. For eg. int8 has an 8 bit size whereas 16 bit has 16bit size

	//	The type of variable can be printed using %T format specifier in Printf function.

	age := 15

	fmt.Printf("Type of number is %T and type of age is %T\n", number, age) // int,int

	//	How to find the size of a variable? --> use Sizeof() method in unsafe package

	fmt.Printf("Size of number is %d\n", unsafe.Sizeof(number))
	fmt.Printf("Size of age is %d\n", unsafe.Sizeof(age))

	//	float type
	marks := 12.3
	fmt.Println(marks)

	//	Complex type
	complexNum := 7 + 5i
	fmt.Println(complexNum)

	// String type
	firstName := "Alex"
	lastName := "Costa"
	fmt.Println(firstName + lastName) // Concatenation will occur(using + operator) since you're adding two strings

	// Type Conversion
	num1 := 12                // int
	num2 := 12.3              // float
	sum := num1 + (int)(num2) // you can't add two numbers of different type, hence type conversion is required
	fmt.Println(sum)
}

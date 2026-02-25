package main

import (
	"fmt"
)

func main() {
	fmt.Println("variables in the golang")

	// using the variable :
	var name string = "aditya"
	fmt.Println(name)

	var curr int = 2434346
	fmt.Println(curr)

	// Here it will automatically detect that it is the string variable
	var myname = "kshatriya"
	fmt.Println(myname)

	var money = 48
	fmt.Println(money)

	fmt.Println("money is", money)

	var dimension float32 = 48.28
	fmt.Println(dimension)

	var istrue bool = false
	fmt.Println(istrue)

	// constant variable declaration in the go lang
	const pi = 3.14
	fmt.Println(pi)

	// shorthand declaration of the variable :
	person := 123
	fmt.Println(person)

	person_name := "aditya kshatriya"
	fmt.Println(person_name)
}

// variable capital intialization --> accessed outside the package
var Public_name = "aditya"

// variable lower case initialization --> only accessed inside the package
var private_name = "kshatriya"

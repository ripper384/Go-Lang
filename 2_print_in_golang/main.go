package main

import (
	"fmt"
)

func main() {
	age := 23
	name := "alice"
	height := 180.38
	salary := 4987990

	//pritln :
	fmt.Println("age: ", age)
	fmt.Println("name : ", name)

	//printf
	fmt.Printf("Name of the person is %s and height of the person is %.4f and the salary is %d\n", name, height, salary)
	// how to find out the datatype of the variable :
	fmt.Printf("datatype of the name is : %T ", name)

}

package myutil

import "fmt"

// function begin with the uppercase letter meant for the exporting and can be used outside the file.
func PrintMessage(message string) {
	fmt.Println(message)
}

// function name begin with the lowercase letter not meant for the exporting and only used within the file
func printMessageAdd(x int, y int) {
	fmt.Println("sum of the x and y : ", x+y)
}

// same is true the variable as well
var name string = "aditya" // cannot be exported
var Value = 2384           // can be exported

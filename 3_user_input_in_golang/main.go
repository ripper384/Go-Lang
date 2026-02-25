// taking user input from the user in the golang
package main

import (
	"fmt"
)


func main(){
	fmt.Println("Enter the name : ") ; 
	var name string ;

	// Scan method reads the input value but it stops reading once it encounters the whitespace
	fmt.Scan(&name) ; 
	fmt.Println("Hello , Mr. " , name) ; 

	// input method that reads the string with the whitespace is : 
	reader:=bufio.NewReader(os.stdin) ; 
	

}
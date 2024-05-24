/*
If Statements involve developing code based on true/false criteria. This controls the programme's flow by branching between different pathways.
*/
/*
This code will output if the integer for the x variable is greater than or less than 10. If it is 10, it will print that it is 10.
*/

package main

import "fmt"

func main() {
	x := 10 /*You can change this to see different outputs in the console*/

	if x > 10 {
		fmt.Println("This number is greater than 10")
	} else if x < 10 {
		fmt.Println("This number is less than 10")
	} else {
		fmt.Println("This number is 10")
	}
}

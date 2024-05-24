/*
In Go, functions are important for creating and organising code that is reusable across your programme. Here is an example of a Go programme that ues functions
*/

package main

import "fmt"

/*Create the function that takes in 2 integers and returns their sum*/
func add(a int, b int) int {
	return a + b
}

/*Create a function that will print a hello message*/
func hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

/*Let's call the function and store results from it*/
func main() {
	sum := add(5, 7)
	fmt.Println("The sum is:", sum)

	/*Let's also call the hello function*/
	hello("Aine")
}
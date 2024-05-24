/*
Arrays are collections of elements that are ordered with a set size. They are accessed by index (starting at 0), and used for storing data types.
*/

package main

import "fmt"

func main() {

	/*Arrays have a length and a datatype*/
	var x [4]int
	fmt.Println(x)

	/*If you want to index an array, you use square brackets. The console output would be [0,0,1,0] here*/
	x[2] = 1
	fmt.Println(x)

	/*If you want to initialise an array, you use curly brackets {}*/
	y := [6]int{7, 6, 5, 4, 3, 2}
	fmt.Println(y)

	/*If you want to create a slice, you remove the element count*/
	c := []int{7, 6, 5, 4, 3, 2}
	fmt.Println(c)

	/*If you want to add something new to the aray, you use append*/
	c = append(c, 1)
	fmt.Println(c)
}

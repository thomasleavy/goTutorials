/*
Loops are useful for recycling code because they allow you to repeat a block of code many times based on a condition.
*/
/*
A for loop is used when the number of iterations is known.
*/
/*
A while loop is used when the condition for it to end is evaluated before each iteration in the loop.
*/
/*
A for each loop is used to iterate over elements within a collection, i.e. arrays and slices.
*/

package main

import "fmt"

func main() {
	/*For loop*/
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	/*While loop*/
	i := 0
	for i < 4 {
		fmt.Println(i)
		i++
	}

	/*For each loop*/
	arr := [4]string{"a", "b", "c", "d"}

	for index, value := range arr {
		fmt.Println("the index is:", index, "the value is:", value)
	}

	m := make(map[string]string)
	m["a"] = "aine"
	m["b"] = "barry"
	m["c"] = "carrie"
	m["d"] = "douglas"

	for key, value := range m {
		fmt.Println("the key is:", key, "the value is:", value)
	}
}

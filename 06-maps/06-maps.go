/*
In Go, maps are collections of key-value pairs. Each key is unique and maps to a particular value. Maps are used to retrieve and store sata in an efficient manner, based on the specified key.
*/

package main

import "fmt"

func main() {
	/*Let's create a map*/
	ages := make(map[string]int)

	/*Let's add elements to the map*/
	ages["Aine"] = 20
	ages["Barry"] = 25
	ages["Carrie"] = 30
	ages["Daithí"] = 35

	/*Let's print elements from the map we created*/
	fmt.Println(ages)
	/*Accessing a key that doesn't exist*/
	fmt.Println(ages["Elsie"])
	/*Delete the entry for Barry*/
	delete(ages, "Barry")
	/*Print the map after Barry was deleted*/
	fmt.Println(ages)
}

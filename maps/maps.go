package main

import (
	"fmt"
)

func main() {
	myMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	//print key and values of map to prove the maps are unordered lists
	for key, value := range myMap {
		fmt.Printf("\nKey is: %v, Value is: %v", key, value)
	}

	//print an individual item of the map
	fmt.Println("\nThe value of the key \"C\" is:", myMap["C"])

	//change the value of an item in the map
	myMap["C"] = 300

	//print the new value in the map
	fmt.Println("\nThe new value of \"C\" is", myMap["C"])

	//Add a new key and value to the map
	myMap["Z"] = 150
	//print the whole map
	fmt.Println(myMap)

	//delete from the map
	delete(myMap, "Z")
	fmt.Println(myMap)
}

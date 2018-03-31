package main

import (
	"fmt"
)

func main() {
	//initialise an array
	myArray1 := [...]int{12, 2, 3, 4}

	// print out array values with for loop
	for i := range myArray1 {
		fmt.Println(myArray1[i])
	}

	//change a value of an array
	myArray1[2] = 10

	// print out array values after update
	fmt.Println(myArray1)

	//slice, initailised 3, length of 5
	mySlice1 := make([]int, 3, 5)
	// print out the slice (will be 3 values of 0)
	fmt.Println(mySlice1)

	//create a slice literal
	mySlice2 := []string{"red", "blue", "green", "yellow"}
	mySlice3 := []int{10, 20, 30}
	fmt.Println(mySlice2)
	fmt.Println(mySlice3)

}

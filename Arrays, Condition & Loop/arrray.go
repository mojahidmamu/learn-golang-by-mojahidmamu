package main

import "fmt"

func main() {
	// Start your code....

	// Array defination:
	//    => Arrays are used to store multiple values of the same type in a single variable,
	//       instead of declaring separate variables for each value.

	/*
	   Syntax :
         var array_name = [array_length] data-type{values}
	*/

	var arr1 = [3]int{1, 2, 3} // data-type : int 
	arr2 := [5]int{1, 2, 3, 5, 8}
   var cars = [4] string{"Volvo", "BMW", "Ford", "Mazda"}; // data-type : string

	fmt.Println(arr1)
	fmt.Println(arr2)
   fmt.Println(cars);
}

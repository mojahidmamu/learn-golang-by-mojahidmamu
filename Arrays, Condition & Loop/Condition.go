package main

import "fmt"

func main() {
	// Start your code....

	// if-else condition:
	// var time = 22
	// if time < 10 {
	// 	fmt.Println("Good morning.")
	// } else if time < 20 {
	// 	fmt.Println("Good day.")
	// } else {
	// 	fmt.Println("Good evening.")
	// }

	// nested if-else conditoin:
	var num = 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}

}

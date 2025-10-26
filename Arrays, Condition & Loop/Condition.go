package main

import "fmt"

func main() {
	// Start your code....
	var time = 22
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
}

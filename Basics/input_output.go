package main

import "fmt"

func main() {
	// Start your code....

	var name string
	var age int

	fmt.Print("Enter you name: ")
	fmt.Scanln(&name) //taking string input from user:

	fmt.Print("Enter you age: ")
	fmt.Scanln(&age) //taking int input from user:

	fmt.Println("Hello ", name)
	fmt.Println("You are ", age, " years old")

}

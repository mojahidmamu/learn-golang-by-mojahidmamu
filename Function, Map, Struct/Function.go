package main

import "fmt"

func message() {
	fmt.Println("Abdullah all MOjahid")
}

func familyName(siblingName string, age int) {
	fmt.Println("Hello" , siblingName, "bin Yousouf and your are", age, "Years old" )
}

func main() {
	// Start your code....
	message() // call the function ...
	message()
	message()

	// 
	familyName("Salma Tasnim", 20)
	familyName("Abdullah all Mojahid",18)
	familyName("Halimatus Sadia",10)
}

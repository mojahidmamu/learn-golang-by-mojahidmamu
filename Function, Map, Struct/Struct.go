package main

import "fmt"

type Person struct { // Struct definition
	Name  string
	Age   int
	Email string
}

func main() {
	// Start your code....
	p1 := Person{Name: "Mojahid", Age: 24, Email: "abdullah@example.com"} // Struct initialization
	fmt.Println(p1)
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Println(p1.Email)
}

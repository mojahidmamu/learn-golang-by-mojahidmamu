package main

import "fmt"

type Person struct { // Struct definition
	Name string
	Age  int
}

func (p Person) Greet() { // Receiver function
	fmt.Printf("Hello, my name is %s. I am %d years old. \n", p.Name, p.Age)
}

func main() {
	// Start your code....
	p1 := Person{Name: "Mojahid", Age: 24} // Struct initialization
	p1.Greet()
	// fmt.Println(p1)
	// fmt.Println(p1.Name)
	// fmt.Println(p1.Age)

}

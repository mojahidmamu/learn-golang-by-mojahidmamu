package main

import "fmt"

func a() {
	i := 0
	defer fmt.Println("defer:", i) // defer statement captures the current value of i = 0
	i++
	defer fmt.Println("defer:", i) // defer statement captures the current value of i = 0

	fmt.Println("i:", i) // prints i: 1

	for i := 1; i <= 3; i++ {
		defer fmt.Println("defer in loop:", i) // defer statements capture i = 1, 2, 3
	}

}

func main() {
	// Start your code....
	a()
}

package main

import "fmt"

func a() {
	i := 0
	defer fmt.Println("defer:", i) // defer statement captures the current value of i = 0
	i++
	fmt.Println("i:", i) // prints i: 1

}

func main() {
	// Start your code....
	a()
}

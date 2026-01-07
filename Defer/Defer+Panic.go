package main

import "fmt"

func main() {
	// Start your code....

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
}

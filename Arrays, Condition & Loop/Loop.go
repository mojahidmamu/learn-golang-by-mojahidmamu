package main

import "fmt"

func main() {
	// Start your code....
	for i := 0; i < 5; i++ {
		if i == 3 {
			break;
		}
		fmt.Println(i)
	}

	for i := 0; i <= 100; i += 10 {
		if i == 50 {
			continue
		}
		fmt.Println(i)
	}
}

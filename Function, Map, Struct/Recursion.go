package main

import "fmt"

func testCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testCount(x + 1)
}

func main() {
	// Start your code....
	testCount(1)
}

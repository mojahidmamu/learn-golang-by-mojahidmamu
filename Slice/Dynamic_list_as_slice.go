package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Start your code....
	var numbers []int

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter numbers (type 'end' to stop)")

	for {
		fmt.Println("> ")
		if !scanner.Scan() {
			break
		}

		text := scanner.Text()
		if text == "end" {
			break
		}

		var x int
		_, err := fmt.Sscanf(text, "%d", &x)

		if err == nil {
			fmt.Println("")
			continue
		}

		numbers = append(numbers, x)
	}
	fmt.Println("You entered :", numbers)

}

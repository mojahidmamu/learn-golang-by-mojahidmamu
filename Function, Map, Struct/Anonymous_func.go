package main
import "fmt"

func main() {
   // Start your code....
    greet := func (name string)  { // Anonymous 
		fmt.Println("Hello" , name)
	}

	greet("Abdullah all Mojahid") // Immediately Invoked Function Expression(IIFE)
}

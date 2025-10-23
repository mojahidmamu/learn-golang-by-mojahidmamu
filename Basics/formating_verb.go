package main

import "fmt"

func main() {
	// Start your code....

	var i = 15.5
	var txt = "Hello world"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	var age = 15

	fmt.Printf("%b\n", age)
	fmt.Printf("%d\n", age)
	fmt.Printf("%+d\n", age)
	fmt.Printf("%o\n", age)
	fmt.Printf("%O\n", age)
	fmt.Printf("%x\n", age)
	fmt.Printf("%X\n", age)
	fmt.Printf("%#x\n", age)
	fmt.Printf("%4d\n", age)
	fmt.Printf("%-4d\n", age)
	fmt.Printf("%04d\n", age)

}

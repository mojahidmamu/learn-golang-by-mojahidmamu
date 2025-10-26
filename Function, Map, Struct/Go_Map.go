package main

import "fmt"

func main() {
	// Start your code....
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Println(a)

	a["year"] = "2025"   // update element
	a["color"] = "Black" // adding element
	fmt.Println(a)

	delete(a, "year") // remove year from map
	fmt.Println(a)
}

package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	// Start your code....
	var per1 Person
	var per2 Person

	// pers1 spacification:
	per1.name = "Mojahid"
	per1.age = 18
	per1.job = "Student"
	per1.salary = 10

	// per2 specification :
	per2.name = "Mamu"
	per2.age = 28
	per2.job = "CNG Driver"
	per2.salary = 10000

	// Access and print Pers1 info
	fmt.Println("Name: ", per1.name)
	fmt.Println("Age: ", per1.age)
	fmt.Println("Job: ", per1.job)
	fmt.Println("Salary: ", per1.salary)

	// Access and print Pers2 info
	fmt.Println("Name: ", per2.name)
	fmt.Println("Age: ", per2.age)
	fmt.Println("Job: ", per2.job)
	fmt.Println("Salary: ", per2.salary)
}

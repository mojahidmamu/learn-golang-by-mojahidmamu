package main
import "fmt"

type Person struct {
	Name string
	Age  int
	Email string
}

func main() {
   // Start your code....
    p1 := Person{Name: "Abdullah all Mojahid", Age: 24, Email: "abdullah@example.com"}
   fmt.Println(p1)
   fmt.Println(p1.Name)
   fmt.Println(p1.Age)
   fmt.Println(p1.Email)
}

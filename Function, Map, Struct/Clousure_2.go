package main
import "fmt"

const a = 10
 
var b = 20

func outer() func() {
	money := 100
	age := 25
	fmt.Println("Age is: " , age)

	show := func() {
		money = a + b
		fmt.Println("Money is: ", money)
	}

	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
   // Start your code....
   fmt.Println("Hello World!")
   call()
}

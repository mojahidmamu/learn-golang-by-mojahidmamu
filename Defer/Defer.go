package main
import "fmt"

func a () {
	i := 0
	defer fmt.Println("defer:", i)
	i++
	fmt.Println("i:", i)

}

func main() {
   // Start your code....
   a()
}

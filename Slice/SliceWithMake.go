package main
import "fmt"

func main() {
   // Start your code....
   arr := [5]int{10,20,30,40,50}

   s1 := arr[1:3]
   fmt.Println("S1 = " , s1)
   fmt.Println("S1 = " , s1[len(s1)-1])
   fmt.Println("Len of s1 = " , len(s1))
   fmt.Println("cap of s1 = " , cap(s1))
}

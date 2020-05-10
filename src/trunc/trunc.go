 package main
 import "fmt"
 
 func main() {
     var num float32
     fmt.Print("Enter the Floating point number: ")
     fmt.Scan(&num)
     ans := int(num)
     fmt.Printf("Truncated number is : %d \n", ans)
 }

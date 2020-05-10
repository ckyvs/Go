 
package main

import (
	"fmt"
	"strings"
)

func main() {
	var st string
	fmt.Print("Enter a string : ")
	fmt.Scan(&st)
	len := len(st)
	str := strings.ToLower(st)
	c := 0
	if str[0] == 'i' && str[len-1] == 'n' {
		if strings.Contains(str, "a") {
			fmt.Print("True\n")
			c = 1
		}
	}
	if c == 0 {
		fmt.Print("False\n")
	}
}

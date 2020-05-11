 package main

import (
	"fmt"
    "bufio"
	"strings"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
	var st string
	fmt.Print("Enter a string : ")
	if scanner.Scan() {
        st = scanner.Text()
    }
	len := len(st)
	str := strings.ToLower(st)
	c := 0
	if str[0] == 'i' && str[len-1] == 'n' {
		if strings.Contains(str, "a") {
			fmt.Print("Found\n")
			c = 1
		}
	}
	if c == 0 {
		fmt.Print("Not Found\n")
	}
}

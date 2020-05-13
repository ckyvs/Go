package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	ar := [3]int{}
	s1 := ar[0:0]
	var x string
	fmt.Scan(&x)
	for x != "X" {
		y, err := strconv.Atoi(x)
		if err == nil {
			s1 = append(s1, y)
			sort.Ints(s1)
		}
		fmt.Println(s1)
		fmt.Scan(&x)
	}
}

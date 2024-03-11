package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 2, 3}
	slices.Reverse(s)
	fmt.Println(s)
}

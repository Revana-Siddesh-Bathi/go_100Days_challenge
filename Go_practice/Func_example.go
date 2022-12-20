package main

import (
	"fmt"
)

func sum(a int, b int) int {
	sumofNUm := a + b
	return sumofNUm
}

func main() {
	Total := sum(3, 7)
	fmt.Println(Total)
}

package main

import "fmt"

var s string

func main() {
	fmt.Println("Enter your name")
	fmt.Scanf("%s", &s)
	fmt.Println("You have entered", s)
}

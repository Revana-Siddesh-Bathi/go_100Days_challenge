// Day 1 challenge
package main

import "fmt"

func main() {
	fmt.Println("******* Calculator app *******")
	var a, b, sum, sub, mul, div int64
	fmt.Println("Enter 2 numbers")
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b
	fmt.Println("You have entered ", a, b)
	fmt.Println("Sum = ", sum)
	fmt.Println("Sub = ", sub)
	fmt.Println("Mul = ", mul)
	fmt.Println("Div = ", div)
}

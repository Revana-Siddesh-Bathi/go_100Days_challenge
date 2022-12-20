package main

import "fmt"

func main() {
	var fruits [5]string = [5]string{"Apple", "mango", "Orange"}
	fmt.Println(fruits)

	names := [4]string{"Sid", "Ranju", "Lasya", "Kriti"}
	fmt.Println(names)

	numbers := [3]int{10, 20, 30}
	fmt.Println(numbers)
	fmt.Println("Number of itmes in Number array is ", len(numbers))
	fmt.Println("3rd Item in the Fruit arrays is ", fruits[2])

	//Looping through the array
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	// looping through range expression
	for index, element := range fruits {
		fmt.Println(index, "=>", element)
	}

	//Multi -Dimensional array -- it is like array of array
	mult_dimen_array := [3][2]int{{2, 3}, {4, 5}, {4, 7}}
	fmt.Println("I am in Mult_Dimen_array")
	fmt.Println(mult_dimen_array[2][1])
}

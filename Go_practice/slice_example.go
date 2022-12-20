// Slice is a continoues segment of an underlying array.
/* Has 3 major components. --
 Pointer -- A variable which holds the address of an array.
 length -- length of the array.
  capacity -- Capacity to hold the elements of the array.
Declaring the array will be same but no need to give any size.

array [start_index : End_index]
*/

package main

import "fmt"

func main() {

	arr := [10]int{10, 20, 30, 20, 30, 50, 40, 70, 90, 00}
	slice_1 := arr[1:9]
	fmt.Println(slice_1)
	fmt.Println(len(slice_1))
	fmt.Println(cap(slice_1))
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

}

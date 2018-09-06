package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3} // array
	slice1 := []int{1, 2, 3}  // slice
	fmt.Println(array1)
	fmt.Println(slice1)

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice define um pedaço de um array
	s2 := a2[1:3] // like python
	fmt.Println(a2, s2)

	s3 := a2[:2]
	fmt.Println(a2, s3)
	fmt.Println(a2[:4])
}

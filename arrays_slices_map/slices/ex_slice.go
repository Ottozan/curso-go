package main

import "fmt"

func main() {

	arr1 := [5]int{1, 2, 3, 4, 5}
	slice1 := arr1[1:3]
	fmt.Println(arr1)
	fmt.Println(slice1)
	arr1[1] = 7
	fmt.Println(slice1, "modificado")

}

package main

import "fmt"

func main() {
	i := 3
	i += 3
	i -= 3
	i /= 2
	i *= 2
	i %= 2
	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}

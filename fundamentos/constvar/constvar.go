package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo inferido float64
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área é ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = a
		d = a + b
	)

	var e, f bool = true, false
	fmt.Println(a, b, c, d, e, f)

	g, h, i := 2, false, "epa"
	fmt.Println(g, h, i)
}

package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é do tipo", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)

	var i2 rune = 'a' // tabela unicode
	fmt.Println(i2)

	s2 := `String
	de
	varias
	linhas`
	fmt.Println("String com linhas", s2)
}

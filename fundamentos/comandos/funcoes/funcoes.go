package main

import (
	"fmt"
)

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}
func main() {
	fmt.Println(somar(3, 5))
	imprimir(123)
}

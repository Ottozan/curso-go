package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
	// uma vez que você atribuiu os valores aos retornos
	// nomeados, não é necessário especificar no return
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)
}

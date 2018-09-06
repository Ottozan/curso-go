package main

import "fmt"

func incrementar1(n int) {
	n++
}

// um ponteiro é representado por um *
func incrementar2(n *int) {
	// o * é usado para desreferenciar, ou seja,
	// ter acesso ao valor para o qual o ponteiro aponta
	*n++
}

func main() {
	n := 1
	incrementar1(n) // passagem por valor
	fmt.Println(n)

	// & é usado pra obter o endereço de uma variável
	incrementar2(&n) // passagem por referência
	fmt.Println(n)
}

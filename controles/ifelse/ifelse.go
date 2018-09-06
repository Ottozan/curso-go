package main

import "fmt"

func imprimirResultado(nota float64) {
	// expressão não tem parênteses, bloco "{}" obrigatório
	if nota >= 7 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
func main() {
	imprimirResultado(8.0)
	imprimirResultado(5.0)
}

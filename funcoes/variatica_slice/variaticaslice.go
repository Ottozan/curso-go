package main

import "fmt"

// Importante: não funciona com arrays, somente slices
func listaAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}
func main() {
	// Colchetes vazios definem var como slice
	aprovados := []string{"Marcos", "João", "Maria", "Guilherme"}

	listaAprovados(aprovados...)
}

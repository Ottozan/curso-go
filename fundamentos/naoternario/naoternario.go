package main

import "fmt"

func obterResultado(nota float64) string {
	// Não existe op ternario em Go
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}
func main() {
	fmt.Println(obterResultado(7.1))
}

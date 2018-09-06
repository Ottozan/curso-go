package main

import "fmt"

func obterValorAprovado(numero int) int {
	// Executa o comando antes do retorno da funçãop
	// normalmente usado para fechar um recurso antes de sair da f
	defer fmt.Println("Fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	// não é necessário o Else, uma vez que se a condição
	// anterior for antendida, há um return
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(6000))
}

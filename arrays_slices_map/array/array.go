package main

import "fmt"

func main() {
	// homogênea (mesmo tipo), estática
	var notas [3]float64
	fmt.Println(notas) // zeros

	notas[0], notas[2], notas[1] = 7.0, 4.3, 9.1
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média de notas: %.2f\n", media)

	// Nro de elementos definidos na inicialização
	numeros := [...]int{1, 2, 3, 4, 5}

	// nesse caso, é obrigado o uso do i (var)
	for i, numero := range numeros {
		fmt.Println(i, ")", numero)
	}

	// Para ignorar o índice, use o underline
	// se informar somente um elemento, ele pega o índice
	for _, numero := range numeros {
		fmt.Printf("Quadrado de %d é %d\n", numero, numero*numero)
	}

}

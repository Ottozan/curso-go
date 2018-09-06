package main

import "fmt"

// número de parâmetros indefinido
func media(numeros ...float64) float64 {
	total := 0.0
	// Parâmetros tratados como um array
	for _, num := range numeros { // _ ignora índice
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.1, 8.5, 9.7, 5.9))

}

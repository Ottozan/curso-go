package main

import "fmt"

func conceito(n float64) string {
	var nota = int(n)
	// note que não precisa do break
	switch nota {
	case 10:
		fallthrough // não faça nada, siga...
	case 9:
		return "A"
	// vários valores: separados por vírgula
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}
func main() {
	fmt.Println("Conceito: ", conceito(10))
	fmt.Println("Conceito: ", conceito(8))
	fmt.Println("Conceito: ", conceito(7))
	fmt.Println("Conceito: ", conceito(5))
	fmt.Println("Conceito: ", conceito(2))
	fmt.Println("Conceito: ", conceito(12))
}

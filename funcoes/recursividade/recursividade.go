package main

import "fmt"

// Dois retornos, um sendo erro
func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

// Usando um meio mais simples só aceita inteiro sem sinal
func fatorial2(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * fatorial2(n-1)

}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	resultSimples := fatorial2(5)
	fmt.Println(resultSimples)
}

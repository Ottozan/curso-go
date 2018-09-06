package main

import "fmt"

func main() {

	funcSalarios := map[string]float64{
		"José":  15000.45,
		"Maria": 16550.00,
		"João":  1200.45, // vircula ou fecha chaves nessa linha
	}

	funcSalarios["Pedro"] = 3450.00

	for nome, salario := range funcSalarios {
		fmt.Printf("%s %.2f)\n", nome, salario)
	}
}

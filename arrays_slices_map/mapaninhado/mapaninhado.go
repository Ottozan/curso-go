package main

import "fmt"

func main() {

	funcPorLetra := map[string]map[string]float64{
		"G": {"Giuliana": 7800.00,
			"Guga": 200.00,
		},
		"J": {
			"José": 8000.00,
			"João": 6000.00,
		},
		"P": {
			"Pedro": 1500.00,
		},
	}

	delete(funcPorLetra, "P")

	for letra, funcionario := range funcPorLetra {
		fmt.Println("Letra", letra)
		for nome, salario := range funcionario {
			fmt.Println(nome, "recebe R$", salario)
		}
	}

}

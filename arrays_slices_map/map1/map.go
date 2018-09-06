package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[48594156987] = "Antonio"
	aprovados[11111111111] = "Jose"
	aprovados[22222222222] = "Maria"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF %d)\n", nome, cpf)
	}
	fmt.Println(aprovados[48594156987])
	delete(aprovados, 48594156987)
	fmt.Println(aprovados)
}

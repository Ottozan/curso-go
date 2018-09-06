package main

import "fmt"

// Nesse caso, para executar, deve-se estar na pasta
// e digitar o comando "go run *.go"

func main() {
	// struct ponto entre {}
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	// A função "cateto" é privada, dentro do pagcote
	// embora esteja em um arquivo separado

	fmt.Println(catetos(p1, p2))
}

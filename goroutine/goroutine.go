package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// Processo em série
	// fale("Maria", "Porque você não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	// go fale("Maria", "Olá", 500)
	// go fale("João", "Fala mulé!", 500)

	// Embora a Maria pretenda falar dez vezes, quando João terminar
	// de falar, a maria também para, porque a rotina main findou.
	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5)
}

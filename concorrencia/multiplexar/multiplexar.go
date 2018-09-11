package main

import (
	"fmt"

	"github.com/ottozan/html"
)

// Primeiro parâmetro, canal somente leitura
func encaminhar(origem <-chan string, destino chan string) {
	// Enquanto tiver dados em origem, manda para canal destino
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.amazon.com", "https://youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}

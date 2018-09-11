package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(3 * time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c) // async
	fmt.Println("A")             // executa imediatamente
	a, b := <-c, <-c             // recebendo dados do canal - sincroniza
	fmt.Println("B")             // aguarda sincronizar *concluir rotinas
	fmt.Println(a, b)

	fmt.Println(<-c)
}

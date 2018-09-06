package main

import "fmt"

// Mesmo sem ser invocada, é a primeira a ser executada
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Método main")
}

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	fmt.Println("Like while...")
	for i < 5 {
		fmt.Print(i, " ")
		i++
	}
	
	fmt.Println("Like true For...")
	for i := 0; i <= 20; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println("")
	// Infinito (Sair com Ctrl+Alt+M)
	for {
		fmt.Println("Pra sempre... ")
		time.Sleep(time.Second)
	}
}

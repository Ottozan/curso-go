package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "desconhecido"
	}
}

func main() {
	fmt.Println(tipo(2.4))
	fmt.Println(tipo(2))
	fmt.Println(tipo("Teste"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}

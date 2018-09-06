package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i // pega endereço de variavel
	*p++   // desreferenciando
	i++
	fmt.Println(p, *p, i, &i)
	// Go não tem aritmética de ponteiros
	// p++

}

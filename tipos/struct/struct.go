package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lápis",
		preco:    1.79,
		desconto: 0.05, // taxa
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Carro", 22000, 0.10}
	fmt.Println(produto2.precoComDesconto())
	produto2.nome = "Chevrolet"
	fmt.Println(produto2.nome, produto2.preco)
}

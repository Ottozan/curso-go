package main

import "fmt"

type item struct {
	produtoID int
	qtdade    int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item // slice do tipo item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtdade)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1, // nomeados, fácil leitura
		itens: []item{
			item{1, 2, 12.10},   // posicional: ID, Qtd, Preço
			item{2, 1, 23.49},   // não aconselhável, pode causar
			item{11, 100, 3.13}, // confusão
		},
	}
	fmt.Printf("Valor total do pedido é %.2f", pedido.valorTotal())
}

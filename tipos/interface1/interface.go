package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	// Sprintf é usado para retornos ou variáveis ...
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Giselda", "Rosa"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// Polimorfismo... coisa agora recebe produot
	coisa = produto{"Calça Jeans", 79.00}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Calça Jeans de Marca", 189.00}
	imprimir(p2)

}

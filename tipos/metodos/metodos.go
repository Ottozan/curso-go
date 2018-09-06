package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

// func (receiver) nomeFuncao() tipoRetorno
func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// Note que a passagem é por ponteiro e não por valor (*)
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1] // só funciona com nome e sobrenome... sem nome do meio
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Maria Braga")
	fmt.Println(p1.getNomeCompleto())
}

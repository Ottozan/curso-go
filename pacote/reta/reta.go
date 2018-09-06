package main

import "math"

// Inicializando ccom letra maiúscula é PÚBLICO (visível dentro e fora do pacote)
// Na linguagem Go, não existe visibilidade a nivel de arquivo, e sim
// dentro de pacotes

// Iniciando com letra minúscula é PRIVADO (dentro do pacote)

// Exemplo:
// Ponto - gerará algo público
// ponto ou _Ponto - privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return // retorno nomeado...
}

// Distancia é responsável por calcular a distância entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

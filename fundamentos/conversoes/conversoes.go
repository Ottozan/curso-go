package main

import 	(
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Conversões")
	x := 2.4
	y := 2
    // Conversão necessária para executar a operação
	fmt.Println(x / float64(y))

	nota := 6.9
	fmt.Println(int(nota))

	// cuidado... 
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste "+strconv.Itoa(57))

	// String para inteiro
	num, _ := strconv.Atoi("123")
	fmt.Println(num)
}

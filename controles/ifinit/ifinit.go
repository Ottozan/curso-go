package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}
func main() {
	// i só é visível dentro da estrutura if
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou", i)
	} else {
		fmt.Println("Perdeu", i)
	}
}
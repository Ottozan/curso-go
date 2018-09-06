package main

import (
	"fmt"
	"time"
)

func main() {
	// ==, !=, >, <, >=, <=

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2, d1, d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Antonio"}
	p2 := Pessoa{"Giuliana"}

	// Compara valores e não endereços de memória
	fmt.Println("Pessoas:", p1 == p2)
}

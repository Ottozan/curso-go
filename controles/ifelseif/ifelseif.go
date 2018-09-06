package main

import "fmt"

func conceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota < 9 {
		return "B"
	} else if nota >= 5 && nota < 8 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	} else {
		return "E"
	}
}
func main() {
	fmt.Println(conceito(9.1))
	fmt.Println(conceito(8.0))
	fmt.Println(conceito(5.0))
	fmt.Println(conceito(3.0))
	fmt.Println(conceito(1.0))
}

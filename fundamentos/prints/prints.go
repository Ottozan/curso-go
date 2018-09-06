package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha")
	fmt.Println(" Nova linha (CR/LF) no fim da linha")
	fmt.Println("Comprovando")

	x := 3.141516
	xs := fmt.Sprint(x) // retorna uma string
	fmt.Println("O valor de x é " + xs)
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "Opa"
	fmt.Printf("\n%d, %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)
}
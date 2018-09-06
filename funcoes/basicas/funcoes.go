package main

import "fmt"

func main() {
	f1()
	f2("Olá", "Mundo")
	r3, r4 := f3(), f4("Sei", "lá")
	fmt.Println(r3, r4)
	r5a, r5b := f5()
	fmt.Println("Retorno F5", r5a, r5b)
}

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s, %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}
package main

//Keywords palavras reservadas não podem ser nomes de variaveis
//statement(declaração, afirmação), uma instrução que forma uma amção composta de expressões
//expressão - qualwuer coisa que produz um resultado
import "fmt"

//func main() {
	//declaração, só funciona dentro de um codeblock,scope
	w := 10.0
	y := "Olá bom dia"

	fmt.Printf("1x: %v, %T\n", w, w)
	fmt.Printf("y: %v, %T\n", y, y)
	//Atribuição
	x := 20
	fmt.Printf("2x: %v, %T\n", x, x)
	//Atribuição com declaração, precisa delcarar pelo menos uma variavel
	x, z := 30, 40
	fmt.Printf("3x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
}

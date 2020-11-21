package main

import "fmt"

//func main() {
	//Cada caractere da String se chama rune
	//strings interpreted literals

	x := "oi bom dia \ncomo vai?\tespero \"que\" tudo bem."
	fmt.Println("Strings interpreted literals =\n", x)

	y := `"oi bom dia \ncomo vai?\tespero \"que\" tudo bem."`
	fmt.Println("Raw string literals= \n", y)

	//Print somente, não adiciona uma linha no final
	a := "Queijo"
	fmt.Print("Terceiro Teste= \n", a, "\n")

	//O Sprint junta duas variáveis, as aspas entre as variaveis serve para dar espaçamento
	variavel1 := "oi"
	variavel2 := "bom dia"
	z := fmt.Sprint(variavel1, " ", variavel2)
	fmt.Print(z)
}

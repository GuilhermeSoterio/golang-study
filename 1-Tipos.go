package main

import "fmt"

//Criado um tipo hotdog com tipo adjacente Int. Porém não da pra relacionar os dois tipos
type hotdog int

var b hotdog = 101

func main() {
	//Criando um tipo hot dog, com um tipo subjacente, o tipo base
	x := 10
	fmt.Printf("z: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", b, b)

	//convertendo o tipo hotdog em inteiro.
	x = int(b)
	fmt.Printf("%v\n", x)
}

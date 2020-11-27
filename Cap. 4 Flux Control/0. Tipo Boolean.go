package main

//- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
import "fmt"

var x bool

//func main() {

	fmt.Println(x) //zero value
	x = true
	fmt.Println(x) //valor atribuido
	x := 10 < 100
	y := 10 == 100
	z := 10 > 100
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

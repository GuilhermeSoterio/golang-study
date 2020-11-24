package main

//- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
import "fmt"

func main() {

	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mês:", mes)
		for x := 1; x <= 30; x++ {
			fmt.Print("Dia: ", x, ", ")
		}
		fmt.Println("\n")

	}
}

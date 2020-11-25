package main

//- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
import "fmt"

func main() {

	x := 0
	//Realizando while
	for {
		if x < 10 {
			x++
			fmt.Println("chis é menor que deis", x)
		} else {
			fmt.Println("chis é maior que deis to fora")
			break
		}
	}

}

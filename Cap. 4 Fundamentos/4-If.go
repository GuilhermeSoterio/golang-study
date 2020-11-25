package main

//- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
import "fmt"

func main() {
	if x := 10; !(x > 100) {
		fmt.Println("Hello, playground")
	}

	if x := 10; !(x > 100) {
		fmt.Println("Hello, playground")
	}else if x < 10 {
		fmt.Println("chis é menor que déis")
	}else if x < 20 {
		fmt.Println("chis é menor que vinte")
	}else if x => 32 {
		fmt.Println("chis é menor que déis")
	} else {
		fmt.Println("chis é 31 ")
	}

}

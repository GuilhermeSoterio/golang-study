package main

//- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i == 3 {
			//Continue quebra apenas essa interação, break quebra o loop inteiro
			//faz isso quando o número não é 3
			continue
		}
		fmt.Println(i)
	}

}

package main

//- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
import "fmt"

func main() {

	quemtanoescritoriohoje := "zezinho"

	switch quemtanoescritoriohoje {
	case "zezinho":
		fmt.Println("hoje quem ta no escritório é o zezinho")
		fallthrough //executa a proxima linha de comando caso seja true
	case "luisinho":
		fmt.Println("hoje quem ta no escritório é o luisinho")
	case "marquinhos":
		fmt.Println("hoje quem ta no escritório é o marquinhos")
	case "joana":
		fmt.Println("hoje quem ta no escritório é o joana")
	default:
		fmt.Println("Tá vazio, nada")
	}

}

package main

// Atribua um valor int a uma variável
// Demonstre este valor em decimal, binário e hexadecimal
// Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
// Demonstre esta outra variável em decimal, binário e hexadecimal
import "fmt"

func main() {

	x := 200
	fmt.Printf("%d\t%#x\t%b\n", x, x, x)
	y := x >> 1
	fmt.Printf("%d\t%b\t%#x", y, y, y)

}

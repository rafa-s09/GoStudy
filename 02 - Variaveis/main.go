package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1" // Declaraçao Explicita
	variavel2 := "Variavel 2"           // Declaraçao Implicitaca (inferencia de tipo)
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	const constant1 string = "Contante 1"
	fmt.Println(constant1)

	// Inverte valores
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}

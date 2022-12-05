package main

import "fmt"

func main() {
	// Declaración de constante:
	const pi float64 = 3.14
	const pi2 = 3.1415
	//fmt.Println(pi, pi2)
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Se crea en el momento
	//base := 15
	//var altura int = 14
	//var area int

	//Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// fmt

	// Classic Println
	message := "Hello Bro"
	fmt.Println(message)

	// Printf
	amountCursos := 233

	fmt.Printf("%s te envió %d saludos \n", message, amountCursos)

	// cuando se desconoce el tipo de dato
	fmt.Printf("%v te envió %v saludos \n", message, amountCursos)

	//Sprintf guarda el mensaje, pero no lo muestra en consola
	messageTwo := fmt.Sprintf("%v te envió %v saludos \n", message, amountCursos)
	fmt.Println(messageTwo)

	// Imprimir tipo de dato
	fmt.Printf("tipo de dato message %T \n", message)
	fmt.Printf("cursos: %T \n", amountCursos)

}

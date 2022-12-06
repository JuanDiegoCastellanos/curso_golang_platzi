package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	// for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("----------------")

	// for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter += 1
	}

	//for forever
	counterForerever := 0
	for {
		fmt.Println(counterForerever)
		counterForerever++
	}

	// condicionales
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// With or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	// Convertir texto a numero
	value, err := strconv.Atoi("53")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Value: ", value)

}

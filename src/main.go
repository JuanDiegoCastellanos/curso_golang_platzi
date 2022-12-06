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
	// counterForerever := 0
	// for {
	// 	fmt.Println(counterForerever)
	// 	counterForerever++
	// }

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

	modulo := 5 % 2

	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	valueS := 200
	switch {
	case valueS > 100:
		fmt.Println("Es mayor a 100")
	case valueS < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condition")
	}

	// defer la linea que contiene defer se ejecuta de ultimas
	fmt.Println("Mundo")
	defer fmt.Println("Hola")

	//Arrays

	var miArreglo [10]int

	miArreglo[0] = 10
	miArreglo[1] = 20
	fmt.Println(miArreglo)

	// len --> cuantos elementos hay en el arreglo
	// cap --> capacidad maxima del arreglo

	//Slices
	slice := []int{0, 2, 3, 4, 2, 1, 2, 4}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append -- agregar elemento
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append -- agregar lista
	newSlice := []int{10, 12, 13, 14, 15}

	slice = append(slice, newSlice...)
	fmt.Println(slice)

}

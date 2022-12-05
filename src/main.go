package main

import "fmt"

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

}

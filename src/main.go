package main

import "fmt"

func message(txt string, c chan string) {
	c <- txt
}

func main() {

	c := make(chan string, 2)

	c <- "Juan"
	c <- "Castellanos"

	fmt.Println(len(c), cap(c))

	// close -- cerrar channel
	close(c)

	for msg := range c {
		fmt.Println(msg)
	}

	// Select
	emailOne := make(chan string)
	emailTwo := make(chan string)

	go message("mensaje1", emailOne)
	go message("mensaje2", emailTwo)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-emailOne:
			fmt.Println("Email recibido de email 1 ", m1)
		case m2 := <-emailTwo:
			fmt.Println("Email recibido de email 2 ", m2)
		}
	}

}

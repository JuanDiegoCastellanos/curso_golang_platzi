package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Concurrencia trabaja con multiples cosas al tiempo")
	fmt.Println("Gorutines")

	say("Hello", &wg)
	wg.Add(1)
	go say("World", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	// time.Sleep(time.Second * 1)
}

package main

import (
	"curso_golang_platzi/src/mypackage"
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}
func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {

	fmt.Println("Modificadores de acceso")

	var myCar mypackage.CarPublico
	myCar.Brand = "Renault"
	fmt.Println(myCar.Brand)

	// Punteros = acceso a memoria
	a := 50
	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}

	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)

	// Stringers
	fmt.Println(myPC)
}

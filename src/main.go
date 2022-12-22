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

// INTERFAZ
type figuras2D interface {
	area() float64
}

func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}

type cuadrado struct {
	base float64
}
type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}
func (r rectangulo) area() float64 {
	return r.base * r.altura
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

	//INTERFACES
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista de interfaces
	myInterface := []interface{}{"Hola", 23, 44.32}
	fmt.Println(myInterface...)
}

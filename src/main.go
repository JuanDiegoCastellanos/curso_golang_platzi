package main

import (
	"curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {

	fmt.Println("Modificadores de acceso")

	var myCar mypackage.CarPublico
	myCar.Brand = "Renault"
	fmt.Println(myCar.Brand)
}

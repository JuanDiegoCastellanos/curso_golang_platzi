package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	// manejador de paquetes
	// go get
	// go modules
	// echo framework

	fmt.Println("Instalar libreria")
	fmt.Println("go install -v github.com/labstack/echo@latest")
	fmt.Println("Inicializar module")
	fmt.Println("go mod init github.com/JuanDiegoCastellanos")

	// Instanciar echo
	e := echo.New()

	//Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.Logger.Fatal(e.Start(":1323"))

}

package main

import (
	"fmt"
)

func saludo(nombre string) {
	var sal, message string
	sal = "¡Hola Mundo"
	message = sal + "! " + nombre
	fmt.Println(message, &sal, &message)
}

func otraFuncion() {
	sal := "¡Hola Mundo"
	nombre := "Nachowski el Grande"
	message := sal + "! " + nombre
	fmt.Println(message, &sal, &message)
}

func main() {
	saludo("Nacho")
	otraFuncion()
}

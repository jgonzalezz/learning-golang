package main

import (
	"fmt"
	"strconv" // convertir enteros a texto
)

var numero int              // default 0
var enterosinsigno uint = 5 // no permite negativos
var texto string            // default ""
var Status bool             // default false   //TODO Primer letra en mayuscula es visible desde otros paquetes SCOPE

func main() {
	//Lenguaje tipado, se debe indicar cada variable de que tipo es
	// los tipos de datos tambien son funciones int(... ....)

	numero = -1

	var numero2 int
	numero3 := 4 // asignacion directa el compilador asume el tipo de dato segun el valor asignado
	numero3 = 5
	numero6, numero7, texto := 2, 4, "Hola Mundo"

	//Definicion de variables encadenadas
	var numero4, numero5 int

	var moneda float32 = 2.8123123
	//Casteo
	numero4 = int(moneda) // Convierte a entero y elimina los decimales

	var texto2 = strconv.Itoa(numero4)

	fmt.Println(numero)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numero5)
	fmt.Println(numero6)
	fmt.Println(numero7)
	fmt.Println(Status)
	fmt.Println(texto)
	fmt.Println(enterosinsigno)
	fmt.Println(texto2)

}

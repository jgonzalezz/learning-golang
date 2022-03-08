package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1, numero2 int
var text, text2 string

func main() {
	fmt.Println("Ingrese numero 1: ")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese numero 2: ")
	fmt.Scanln(&numero2)

	fmt.Println("Ingrese la descripcion : ")
	fmt.Scanln(&text)
	fmt.Println(text, numero1+numero2)

	//Nota: utilizar la siguiente configuracion para mejores resultados
	fmt.Println("Ingrese la descripcion 2 : ")
	scanner := bufio.NewScanner(os.Stdin) //Standard In es el teclado del OS
	if scanner.Scan() {
		text2 = scanner.Text()
	}
	fmt.Println(text2, numero1+numero2)

}

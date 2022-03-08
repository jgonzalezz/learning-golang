package main

import "fmt"

func main() {

	fmt.Println(uno(3))
	fmt.Println(unoVariante(3))

	fmt.Println(dos(3))

	numero, estado := dos(4)
	fmt.Println("Variables que reciben los valores directamente de la funcion ", numero, "   ", estado)

	fmt.Println(calculo(34, 23, 45, 455, 34, 3, 23))
	fmt.Println(calculo2(23, 45, 65, 7, 8, 3))
}

//Recibe un entero devuelve un entero
func uno(numero int) int {
	return numero * 2
}

//Variante
func unoVariante(numero int) (z int) {
	z = numero * 2
	return
}

//Recibe un entero devuelve un entero y un booleano

func dos(numero int) (int, bool) {
	return numero * 2, true
}

// ... vas a recibir parametro enteros pero no se cuantos
func calculo(numero ...int) int {
	total := 0

	//debo iterar cada parametro que recibo
	for _, num := range numero { //range convierte el numero en una lista  // el _ se usa para obviar(y alojar) el tipo de dato que no voy a usar
		total = total + num
	}
	return total
}

// ... vas a recibir parametro enteros pero no se cuantos
func calculo2(numero ...int) int {
	total := 0
	for x, num := range numero { // en este caso si utilizo la x
		if x == 2 || x == 5 {
			total = total + num
		}
	}
	return total
}

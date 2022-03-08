package main

import "fmt"

//Funciones Anomina
//No tienen nombre, puedeo crear una variblae de tipo funcion para poder modificarla en tiempo de ejecucion
//Variable global de tipo funcion
var Calculo func(int, int) int // recibe 2 enteros y devuelve un entero

func main() {

	//Sumamos
	Calculo = func(num1 int, num2 int) int {
		return num1 + num2
	}

	fmt.Println("Sumo 7 + 5 = ", Calculo(7, 5))

	//Restamos redifiniendo la funcion anomina Calculo
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Println("Restamos 7 - 5 = ", Calculo(7, 5))

	//Multiplicamos redifiniendo la funcion anomina Calculo
	Calculo = func(num1 int, num2 int) int {
		return num1 * num2
	}

	fmt.Println("Multiplicamos 7 * 5 = ", Calculo(7, 5))

	Operaciones()

	//Closures
	MiTabla := Tabla(2)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}

}

func Operaciones() {
	resultado := func() int {
		var a int = 5
		var b int = 7
		return a + b
	}

	fmt.Println(resultado())
}

//Closures : funciones anominas con codigo y variables seguras (isolacion de codigo) pueden acceder a variables creadas por fuera de la funcion
func Tabla(valor int) func() int {

	//LO MAS IMPORTANTE: Las variables definidas por fuera de la funcion de retorno mantienen su valor en memoria, no cambian.
	numero := valor
	secuencia := 0

	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}

package main

import "fmt"

//Vector
var tabla [10]int

//matriz
var matriz [5][7]int

func main() {

	//En go todas las declaraciones se inicilizan en cero cadena vacia o false por defecto
	tabla[0] = 1
	tabla[5] = 15

	fmt.Println(tabla)

	tabla1 := [10]int{2, 34, 4, 6, 67, 7, 9, 6, 7, 4}
	fmt.Println(tabla1)

	for i := 0; i < len(tabla1); i++ {
		fmt.Println(tabla1[i])
	}

	matriz[3][5] = 1
	fmt.Println(matriz)

	/* Slices: Vectores dinamicos, donde puedeo ampliar las dimenciones en tiempo de ejecucion (parecido a las listas) */

	matriz2 := []int{3, 5, 4} // se declara sin tamaño
	fmt.Println(matriz2)

	elementos := [5]int{4, 5, 6, 7, 8}
	porcion := elementos[3:] //creo un slice a partir de otro vector
	fmt.Println(porcion)
	porcion2 := elementos[:2] //creo un slice a partir de otro vector
	fmt.Println(porcion2)

	elementos2 := make([]int, 5, 20) // capacidad vs largo de un vector... en memoria reserva lo que ponemos de capacidad.. eje: crea una matriz de 5 pero puede crecer hasta 20
	fmt.Printf("Largo %d, capacidad %d ", len(elementos2), cap(elementos2))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i) //agrega elementos al slice redimencionando la matriz en el largo y la capacidad
	}
	fmt.Printf("\n Largo %d, capacidad %d ", len(nums), cap(nums))

}

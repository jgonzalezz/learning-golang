package main

import "fmt"

func main() {

	//TODO ciclos for break, continue y rutinas

	//do while , while --- No existen!!! solo se usa FOR

	fmt.Println("-------FOR----------")

	i := 1
	for i <= 10 {
		fmt.Print(i)
		i++
	}
	fmt.Println()
	for x := 1; x <= 10; x++ {
		fmt.Print(x)
	}

	fmt.Println()
	fmt.Println("-------BREAK----------")

	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 7 {
			break //rompe la interacion
		}
	}

	fmt.Println()
	fmt.Println("-------CONTIUE----------")

	var y = 0
	for y < 10 {
		fmt.Printf("\n valor de y: %d", y)
		if y == 5 {
			fmt.Print(" le sumamos  2 y no debe ejecutar el paso por alla :D \n")
			y = y + 2
			continue // Retorna al inicio de la iteracion sin ejecutar lo que este abajo
		}
		fmt.Println("  Paso por aqui")
		y++
	}

	fmt.Println()
	fmt.Println("-------RUTINA----------")

	var z = 0

RUTINA:
	fmt.Printf("Se devolvio a la rutina \n")

	for z < 10 {
		if z == 4 {
			fmt.Printf("Le sumamos 2 y no debe mostrar el Valor de z: %d\n", z)
			z = z + 2
			goto RUTINA // A diferencia del continue este retorna al inicio de la rutina y no al for
		}
		fmt.Printf("Valor de z: %d\n", z)
		z++
	}

}

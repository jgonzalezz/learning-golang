package main

import "fmt"

var status bool

func main() {

	//IF
	status = true
	if status == true {
		fmt.Println(status)
	} else {
		fmt.Println("Es falso")
	}

	//Asignar valor a una variable en el if
	if status = false; status == true {
		fmt.Println(status)
	} else {
		fmt.Println("Es falso")
	}

	//SWITCH
	switch numero := 5; numero {
	case 1:
		fmt.Println(1)
	case 5:
		fmt.Println("numero:", 5)

		//No hace falta agregar break
	}

}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	archivo := "prueba.txt"
	f, err := os.Open(archivo)

	//DEFER : Instruccion que se ejecuta si o si cuando se detecta se va por un return o por un error o por un fin de funcion
	defer f.Close() // se ejecuta al final

	if err != nil {
		fmt.Println("Error abriendo el archivo")
		//os.Exit(1) // finaliza la ejecucion del programa
	}

	//PANIC: Forzar un error
	ejemploPanic()

}

func ejemploPanic() {

	//RECOVER: para evitar que el programa se cierre con el panic y guardo el log
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que gener un panic \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Panic: encontro el valor de 1")
	}
}

package main

import (
	"fmt"
	"strings"
	"time"
)

//Ejecuciones asinconas

//CHANELS: Canales nos permite que una goRutine envie info hacia otra fucion o haciaa otra gorutine
// Es un espacio de memria de dialogo, donde van a dialogar dstantas rutinas... debe ser de un tipo de dato

//go y chan
func main() {

	canal1 := make(chan time.Duration)
	go bucle(canal1) //ejecuta lo que tenga la rutina bucle pero toma en cuenta que te estoy pasando un canal para dialogar con otras rutinas
	fmt.Println("Llegue hasta aqui")

	msg := <-canal1 // await : hasta que no recibe un valor en el canal
	fmt.Println(msg)

	go miNombreLento("jhonatan")

	//Ejecuto otras tares
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	time.Sleep(time.Second * 5)
	final := time.Now()
	canal1 <- final.Sub(inicio) // diferencia de tiempo y la agrego al canal

}

//Una ejecucion lenta
func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "") //convierto un sting en array
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond) // 1 segundo
		fmt.Println(letra)
	}
}

package main

import (
	"fmt"
	"time"

	us "./user"
)

// ESTRUCTURAS   (Son Muy Importantes en Go)
// NO EXISTE la POO

type usuario1 struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

//Creo una nueva estructura para ver la HERENCIA, quiero HEREDAR de Usuario
type pepe struct {
	us.Usuario        // HERENCIA
	Telefono   string // polimorfismo --Atributo que no esta en el la estructura original Usuario
}

func main() {
	User1 := new(usuario1)
	User1.Id = 10
	User1.Nombre = "Pablo"
	fmt.Println(User1)

	//LLamando(instanciando) estructura de otro paquete user
	u := new(pepe)
	u.Telefono = "2343445"
	u.AltaUsuario(1, "Pablo", time.Now(), true)
	fmt.Println(u.Usuario, u.Telefono, u.Nombre)

	// Sin herencia
	u2 := new(us.Usuario)
	u2.AltaUsuario(2, "Pablo2", time.Now(), false)
	fmt.Println(u2.Id, u2.Nombre, u2.Status)

}

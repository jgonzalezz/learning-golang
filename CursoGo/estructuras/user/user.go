package user

import "time"

type Usuario struct { //Usuario con mayuscula para que tenga total visibilidad fuera de este paquete
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

//Vamos agregar una funcion (Method) que asigna valores como si fuera el costructor
func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}

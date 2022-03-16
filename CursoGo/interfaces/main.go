package main

import "fmt"

// se definen los metodos que vamos usar para implementar la interface
type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
}

type mujer struct {
	hombre // mujer hereda las propiedades de l hombre
}

//Se implementan los metodos para la estructura que va utilizar la interfaz
func (h *hombre) respirar() {
	h.respirando = true
}
func (h *hombre) pensar() { h.comiendo = true }
func (h *hombre) comer()  { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}
}

// en este metodo implemento un metodo de la interfaz
func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

// Utilizando la interfaz animal
type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

// en este metodo implemento un metodo de la interfaz
func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

func main() {

	//Humanos
	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanosRespirando(Pedro)
	fmt.Println("Pedro esta respirando: ", Pedro.respirando)

	Maria := new(mujer)
	HumanosRespirando(Maria)
	fmt.Println("Maria esta respirando: ", Pedro.respirando)

	//Animales
	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	AnimalesRespirar(Dogo)
	totalCarnivoros = +AnimalesCarnivoros(Dogo)

	fmt.Printf("Total carnivoros %d", totalCarnivoros)

}

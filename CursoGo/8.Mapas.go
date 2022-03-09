package main

import "fmt"

func main() {
	//Mapas

	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)

	campeonato := map[string]int{
		"Real Madrid": 30,
	}

	campeonato["Barcelona"] = 39

	fmt.Println(campeonato)

	//Eliminar elemento del mapa
	delete(campeonato, "Barcelona")
	fmt.Println(campeonato)

	//Recorrer mapa con range... No ordena
	for pais, capital := range paises {
		fmt.Println(pais, capital)

	}

	//Preguntar si un elemento existe en el mapa
	capital, existe := paises["Mexico"]
	fmt.Println(capital, existe)

}

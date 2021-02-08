package main

import (
	"./estructura"
	"./tienda"
)

func main() {

	tiendaprueba := tienda.Store{"Hola", "holaputa", "124537", 5}
	tiendaprueba2 := tienda.Store{"Hola2", "holaputito", "124537", 5}
	tiendaprueba3 := tienda.Store{"Tienda1", "hola", "124537", 5}
	tiendaprueba4 := tienda.Store{"Tienda2", "hola", "124537", 5}

	lista := estructura.List{}

	lista.Add(tiendaprueba)
	lista.Add(tiendaprueba2)
	lista.Add(tiendaprueba3)
	lista.Add(tiendaprueba4)
	lista.Printlist()
	lista.DeleteStore("Tienda1")
	lista.Printlist()

}

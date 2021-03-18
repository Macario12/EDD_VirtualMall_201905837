package matrizDispersa

import (
	"fmt"
	"strconv"

	"../pedidos"
)

type nodoCola struct {
	pedido    pedidos.Pedido
	siguiente *nodoCola
}

type Cola struct {
	cabeza   *nodoCola
	final    *nodoCola
	Longitud int
}

func (c *Cola) estaVacia() bool {
	if c.cabeza == nil {
		return true
	}
	return false
}

//Encolar ...
func (c *Cola) Encolar(p pedidos.Pedido) {
	nuevoNodo := &nodoCola{pedido: p}
	if c.estaVacia() {
		c.cabeza = nuevoNodo
		c.final = nuevoNodo
	} else {
		nuevoNodo.siguiente = c.cabeza
		c.cabeza = nuevoNodo
	}
	c.Longitud++
}

//Pop ...
func (c *Cola) Pop() pedidos.Pedido {
	var eliminado = c.final
	if c.cabeza == c.final {
		c.cabeza = nil
		c.final = nil
	} else {
		nodoActual := c.cabeza
		for nodoActual.siguiente != eliminado {
			nodoActual = nodoActual.siguiente
		}
		c.final = nodoActual
		c.final.siguiente = nil
	}
	c.Longitud--
	return eliminado.pedido
}

//Imprimir ...
func (c *Cola) Imprimir() {
	texto := ""
	nodoActual := c.cabeza
	for nodoActual != nil {
		texto += strconv.Itoa(nodoActual.pedido.Calificacion) + " -> "
		if nodoActual.siguiente != nil {
			nodoActual = nodoActual.siguiente
		} else {
			nodoActual = nil
		}
	}
	texto += "nil"
	fmt.Println(texto)
}

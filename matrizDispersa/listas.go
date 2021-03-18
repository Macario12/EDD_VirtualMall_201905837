package matrizDispersa

import (
	"../arbol_avl"
	"../tienda"
)

type nodo struct {
	Valor      interface{}
	siguiente  *nodo
	anterior   *nodo
	Inventario *arbol_avl.Tree
}

func (n *nodo) getSiguiente() *nodo {
	return n.siguiente
}

func (n *nodo) setSiguiente(nodo *nodo) {
	n.siguiente = nodo
}

func (n *nodo) getAnterior() *nodo {
	return n.anterior
}

func (n *nodo) setAnterior(nodo *nodo) {
	n.anterior = nodo
}

// ListaDoble ...
type ListaDoble struct {
	cabeza   *nodo
	cola     *nodo
	longitud int
}

//GetLen ...
func (lista ListaDoble) GetLen() int {
	return lista.longitud
}

//Mostrar ...
func (lista ListaDoble) Mostrar() string {
	if lista.longitud == 0 {
		return "ListaDoble vacía"
	}
	resultado := ""
	nodoActual := lista.cabeza
	for i := 0; i < lista.longitud; i++ {
		resultado += nodoActual.Valor.(tienda.Store).Name + " -> "
		nodoActual = nodoActual.getSiguiente()
	}
	resultado = resultado[:len(resultado)-4]
	return resultado
}

//Agregar ...
func (lista *ListaDoble) Agregar(i interface{}) {
	newNodo := nodo{i, nil, nil, new(arbol_avl.Tree)}
	if lista.cabeza == nil {
		lista.cabeza = &newNodo
		lista.cola = &newNodo
	} else {
		lista.cola.setSiguiente(&newNodo)
		newNodo.setAnterior(lista.cola)
		lista.cola = &newNodo
	}
	lista.longitud++
}

//Eliminar ...
func (lista *ListaDoble) Eliminar(i interface{}) {
	nodoActual := lista.cabeza
	for nodoActual != nil && nodoActual.Valor != i {
		nodoActual = nodoActual.getSiguiente()

	}
	if nodoActual != nil {
		if lista.longitud == 1 {
			lista.cabeza = nil
			lista.cola = nil
		} else if lista.cabeza.Valor == i {
			lista.cabeza = lista.cabeza.getSiguiente()
			lista.cabeza.setAnterior(nil)
		} else if lista.cola.Valor == i {
			lista.cola = lista.cola.getAnterior()
			lista.cola.setSiguiente(nil)
		} else {
			nodoActual.getSiguiente().setAnterior(nodoActual.getAnterior())
			nodoActual.getAnterior().setSiguiente(nodoActual.getSiguiente())
		}
		lista.longitud--
	}
}

// Encontrar ...
func (lista *ListaDoble) Encontrar(i interface{}) int {
	if lista.cabeza == nil {
		return -1
	}
	if lista.cabeza.Valor == i {
		return 0
	} else if lista.cola.Valor == i {
		return lista.longitud - 1
	}
	indice := 0
	nodoActual := lista.cabeza
	for nodoActual.Valor != i && nodoActual.getSiguiente() != nil {
		nodoActual = nodoActual.getSiguiente()
		indice++
	}
	if nodoActual.Valor == i {
		return indice
	}
	return -1
}

//Obtener ...
func (lista *ListaDoble) Obtener(indice int) interface{} {
	if indice >= 0 && indice < lista.longitud {
		nodoActual := lista.cabeza
		for i := 0; i < indice; i++ {
			nodoActual = nodoActual.getSiguiente()
		}
		return nodoActual.Valor
	}
	return nil
}

//ObtenerNodo ...
func (lista *ListaDoble) ObtenerNodo(indice int) *nodo {
	if indice >= 0 && indice < lista.longitud {
		nodoActual := lista.cabeza
		for i := 0; i < indice; i++ {
			nodoActual = nodoActual.getSiguiente()
		}
		return nodoActual
	}
	return nil
}

//Vaciar ...
func (lista *ListaDoble) Vaciar() {
	lista.cabeza = nil
	lista.cola = nil
	lista.longitud = 0
}

//Contiene ...
func (lista *ListaDoble) Contiene(i interface{}) bool {
	for a := 0; a < lista.longitud; a++ {
		var elemento interface{} = lista.Obtener(a)
		if elemento == i {
			return true
		}
	}
	return false
}

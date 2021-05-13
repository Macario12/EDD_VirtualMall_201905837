package estructura

import (
	"fmt"

	"../Tabla_Hash"
	"../tienda"
)

type Nodo struct {
	Store       tienda.Store
	Next        *Nodo
	Previous    *Nodo
	Comentarios *Tabla_Hash.TablaHash
}

func NewNodo(Store tienda.Store) *Nodo {
	return &Nodo{
		Store: Store,
		Comentarios:  Tabla_Hash.NewTablaHash(7),
	}
}

func (n *Nodo) GetTabla() *Tabla_Hash.TablaHash {
	return n.Comentarios
}


func (n *Nodo) SetStore(Store tienda.Store) {
	n.Store = Store
}

func (n *Nodo) SetNext(Next Nodo) {
	n.Next = &Next
}

func (n *Nodo) SetPrevious(Previous Nodo) {
	n.Previous = &Previous
}

func (n Nodo) GetStore() tienda.Store {
	return n.Store
}

func (n Nodo) GetNext() *Nodo {
	return n.Next
}

func (n Nodo) GetPrevious() *Nodo {
	return n.Previous
}

type List struct {
	Frist *Nodo
	Last  *Nodo
	size  int
}

func (s *List) Add(e tienda.Store) {

	nodoaux := NewNodo(e)

	nodoaux.Next = s.Frist
	nodoaux.Previous = nil

	if s.Frist != nil {
		s.Frist.Previous = nodoaux
	}

	s.Frist = nodoaux
	s.size++
}
func (lista List) GetLen() int {
	return lista.size
}
func (s List) Printlist() {
	imprimir := s.Frist
	for s.size != 0 {
		fmt.Println(imprimir.Store.GetName())
		imprimir = imprimir.GetNext()
		s.size--

	}

}

func (s *List) DeleteStore(Nombre string) bool {
	aux := s.Frist

	if s.Frist != nil {
		for aux != nil {
			if aux.Store.Name == Nombre {
				if s.Frist == aux {
					s.Frist = aux.Next
				}

				if aux.Next != nil {
					aux.Next.Previous = aux.Previous
				}

				if aux.Previous != nil {
					aux.Previous.Next = aux.Next
				}

				return true

			}
			aux = aux.Next
		}
	}

	return false
}

func (lista *List) ObtenerNodo(indice int) *Nodo {
	if indice >= 0 && indice < lista.size {
		nodoActual := lista.Frist
		for i := 0; i < indice; i++ {
			nodoActual = nodoActual.GetNext()
		}
		return nodoActual
	}
	return nil
}

func (s *List) SearchStore(Nombre string) tienda.Store {
	tienda := tienda.Store{}
	if s.Frist != nil {
		Nodoaux := s.Frist

		for Nodoaux != nil {
			if Nodoaux.GetStore().GetName() == Nombre {
				tienda = Nodoaux.GetStore()
				return tienda

			}
			Nodoaux = Nodoaux.Next
		}

	}

	return tienda
}

func (s *List) SearchNodo(Nombre string, Contact string) *Nodo {
	var Nodoaux *Nodo
	if s.Frist != nil {
		Nodoaux = s.Frist

		for Nodoaux != nil {
			if Nodoaux.GetStore().GetName() == Nombre && Nodoaux.GetStore().GetContact() == Contact{
				return Nodoaux

			}
			Nodoaux = Nodoaux.Next
		}

	}

	return Nodoaux
}

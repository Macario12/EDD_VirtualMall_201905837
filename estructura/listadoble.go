package estructura

import (
	"fmt"

	"../tienda"
)

type Nodo struct {
	Store    tienda.Store
	Next     *Nodo
	Previous *Nodo
}

func NewNodo(Store tienda.Store) *Nodo {
	return &Nodo{
		Store: Store,
	}
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

	if s.Frist == nil {
		s.Frist = nodoaux
	} else {
		nodoaux.SetNext(*s.Frist)
		nodoaux.SetPrevious(*nodoaux)

		s.Frist = nodoaux
	}
	s.size++
}

func (s List) Printlist() {
	imprimir := s.Frist
	for s.size != 0 {
		fmt.Println(imprimir.Store.GetName())
		imprimir = imprimir.GetNext()
		s.size--

	}

}

func (s *List) DeleteStore(Nombre string) {
	if s.Frist != nil {
		nodoaux := s.Frist
		nodoant := new(Nodo)

		for nodoaux != nil {

			if nodoaux.GetStore().GetName() == Nombre {
				if nodoant == nil {
					s.Frist = s.Frist.GetNext()
					nodoaux.Next = nil
					nodoaux = nodoant.GetNext()
				} else {
					nodoant.SetNext(*nodoaux.GetNext())
					nodoaux.Next = nil
					nodoaux = nodoant.GetNext()
				}
			} else {
				nodoant = nodoaux
				nodoaux = nodoaux.GetNext()
			}

		}
	}
	s.size--
}

func (s *List) SearchStore(Nombre string) {
	if s.Frist != nil {
		Nodoaux := s.Frist

		for Nodoaux != nil {
			if Nodoaux.GetStore().GetName() == Nombre {
				fmt.Println(Nodoaux.GetStore().GetDescription())

			}
			Nodoaux = Nodoaux.Next
		}

	}
}


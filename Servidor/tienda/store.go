package tienda

import "../arbol_avl"

type Store struct {
	Name        string `json:"Nombre"`
	Description string `json:"Descripcion"`
	Contact     string `json:"Contacto"`
	Score       int    `json:"Calificacion"`
	Logo        string `json:"Logo"`
	Productos   *arbol_avl.Tree
	Arbol64     *string
}

func (n *Store) SetName(Name string) {
	n.Name = Name
}

func (n *Store) SetDescription(Description string) {
	n.Description = Description
}

func (n *Store) SetContact(Contact string) {
	n.Contact = Contact
}

func (n *Store) SetScore(Score int) {
	n.Score = Score
}

func (n Store) GetName() string {
	return n.Name
}

func (n Store) GetDescription() string {
	return n.Description
}

func (n Store) GetContact() string {
	return n.Contact
}

func (n Store) GetScore() int {
	return n.Score
}

func (n Store) GetProductos() *arbol_avl.Tree {
	return n.Productos
}

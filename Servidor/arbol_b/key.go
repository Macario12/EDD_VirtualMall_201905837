package arbol_b


type Key struct {
	Value     User
	Izquierdo *Nodo
	Derecho   *Nodo
}

func NewKey(valor User) *Key {
	k := Key{valor, nil, nil}
	return &k
}

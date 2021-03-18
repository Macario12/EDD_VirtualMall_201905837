package arbol_avl

type Inventarioss struct {
	Inventarios []Datos `json:"Inventarios"`
}

type Datos struct {
	Tienda       string       `json:"Tienda"`
	Departamento string       `json:"Departamento"`
	Calificacion int          `json:"Calificacion"`
	Productoss   []Inventario `json:"Productos"`
}

var VectorInventario Inventarioss

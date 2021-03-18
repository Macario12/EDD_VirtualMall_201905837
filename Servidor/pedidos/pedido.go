package pedidos

type PedidosP struct {
	PedidosPr []Pedido `json:"Pedidos"`
}

type Pedido struct {
	Fecha        string      `json:"Fecha"`
	Tienda       string      `json:"Tienda"`
	Departamento string      `json:"Departamento"`
	Calificacion int         `json:"Calificacion"`
	ProductosP   []ProductoP `json:"Productos"`
}

type ProductoP struct {
	Codigo int `json:"Codigo"`
}

var VectorPedidos PedidosP

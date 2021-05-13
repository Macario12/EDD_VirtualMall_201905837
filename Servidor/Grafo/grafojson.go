package Grafo

var Grafojson GrafoArc

type GrafoArc struct {
	Nodos                []Nodo `json:"Nodos"`
	PosicionInicialRobot string `json:"PosicionInicialRobot"`
	Entrega              string `json:"Entrega"`
}

type Nodo struct {
	Nombre    string     `json:"Nombre"`
	EnlancesA []EnlanceA `json:"Enlaces"`
}

type EnlanceA struct {
	Nombre    string `json:"Nombre"`
	Distancia int    `json:"Distancia"`
}

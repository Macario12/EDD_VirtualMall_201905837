package matrizDispersa

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"../pedidos"
)

//MatrisDispersa ...
type MatrisDispersa struct {
	CabezeraFilas    *NodoMatris
	CabezeraColumnas *NodoMatris
}

type NodoMatris struct {
	superior     *NodoMatris
	inferior     *NodoMatris
	posterior    *NodoMatris
	anterior     *NodoMatris
	EsCabezera   bool
	Dia          string
	Departamento string
	cola         Cola
	PosX         int
	PosY         int
}

// ___________________________________ getPuntero ____________________________________
func (n *NodoMatris) getSuperior() *NodoMatris {
	return n.superior
}
func (n *NodoMatris) getInferior() *NodoMatris {
	return n.inferior
}
func (n *NodoMatris) getPosterior() *NodoMatris {
	return n.posterior
}
func (n *NodoMatris) getAnterior() *NodoMatris {
	return n.anterior
}

// ___________________________________ setPuntero ____________________________________
func (n *NodoMatris) setSuperior(nodo *NodoMatris) {
	n.superior = nodo
}
func (n *NodoMatris) setInferior(nodo *NodoMatris) {
	n.inferior = nodo
}
func (n *NodoMatris) setPosterior(nodo *NodoMatris) {
	n.posterior = nodo
}
func (n *NodoMatris) setAnterior(nodo *NodoMatris) {
	n.anterior = nodo
}

// ___________________________________ ObtenerId _______________________________
func (n *NodoMatris) getId() string {
	if n.EsCabezera {
		if n.PosX == -1 {
			return strings.ReplaceAll(n.Departamento, " ", "_")
		} else {
			return n.Dia
		}
	}
	return strconv.Itoa(n.PosX) + "y" + strconv.Itoa(n.PosY)

}

// __________________________________ getEncabezadoVertical ___________________________________
func (m *MatrisDispersa) getEncabezadoVertical(Dia string, aInsertar *NodoMatris) *NodoMatris {
	nodoActual := m.CabezeraColumnas
	for nodoActual.getPosterior() != nil && nodoActual.Dia != Dia {
		nodoActual = nodoActual.getPosterior()
	}
	if nodoActual.Dia != Dia {
		nodoActual = m.CabezeraColumnas
		for nodoActual.getPosterior() != nil && nodoActual.getPosterior().Dia < Dia {
			nodoActual = nodoActual.getPosterior()
		}
		if nodoActual.Dia < Dia {
			if nodoActual.getPosterior() != nil {
				n2 := nodoActual.getPosterior()
				nodoActual.setPosterior(aInsertar)
				n2.setAnterior(aInsertar)
				aInsertar.setAnterior(nodoActual)
				aInsertar.setPosterior(n2)
			} else {
				nodoActual.setPosterior(aInsertar)
				aInsertar.setAnterior(nodoActual)
			}
		} else if nodoActual.Dia > Dia {
			nodoActual.setAnterior(aInsertar)
			aInsertar.setPosterior(nodoActual)
			m.CabezeraColumnas = aInsertar
		}
		nodoActualizar := m.CabezeraColumnas
		contador := 0
		for nodoActualizar != nil {
			nodoActualizar.PosX = contador
			//actualizar para abajo
			nodoActualizar2 := nodoActualizar
			for nodoActualizar2 != nil {
				nodoActualizar2.PosX = contador
				nodoActualizar2 = nodoActualizar2.getInferior()
			}
			contador++
			nodoActualizar = nodoActualizar.getPosterior()
		}
		return aInsertar
	}
	return nodoActual
}

// _________________________________ getEncabezadoHorizontal __________________________________
func (m *MatrisDispersa) getEncabezadoHorizontal(Departamento string, aInsertar *NodoMatris) *NodoMatris {
	nodoActual := m.CabezeraFilas
	for nodoActual.getInferior() != nil && nodoActual.Departamento != Departamento {
		nodoActual = nodoActual.getInferior()
	}

	if nodoActual.Departamento != Departamento {
		nodoActual = m.CabezeraFilas
		for nodoActual.getInferior() != nil && nodoActual.getInferior().Departamento < Departamento {
			nodoActual = nodoActual.getInferior()
		}
		if nodoActual.Departamento < Departamento {
			if nodoActual.getInferior() != nil {
				n2 := nodoActual.getInferior()
				nodoActual.setInferior(aInsertar)
				n2.setSuperior(aInsertar)
				aInsertar.setSuperior(nodoActual)
				aInsertar.setInferior(n2)
			} else {
				nodoActual.setInferior(aInsertar)
				aInsertar.setSuperior(nodoActual)
			}
		} else {
			nodoActual.setSuperior(aInsertar)
			aInsertar.setInferior(nodoActual)
			m.CabezeraFilas = aInsertar
		}
		nodoActualizar := m.CabezeraFilas
		cotnador := 0
		for nodoActualizar != nil {
			nodoActualizar.PosY = cotnador
			//actualizar para la derecha
			nodoActualizar2 := nodoActualizar
			for nodoActualizar2 != nil {
				nodoActualizar2.PosY = cotnador
				nodoActualizar2 = nodoActualizar2.getPosterior()
			}
			cotnador++
			nodoActualizar = nodoActualizar.getInferior()
		}
		return aInsertar
	}
	return nodoActual
}

// ___________________________ ingresa pedido a casilla ya existente ___________________________
func actualizar(columnaN *NodoMatris, filaN *NodoMatris, pedido pedidos.Pedido) bool {
	columna := columnaN
	fila := filaN
	for columna != fila && !(columna.getInferior() == nil && fila.getPosterior() == nil) {
		if columna.getInferior() != nil {
			columna = columna.getInferior()
		}

		if fila.getPosterior() != nil {
			fila = fila.getPosterior()
		}
	}
	if fila == columna {
		fila.cola.Encolar(pedido)
		return true
	}
	return false
}

// _____________________________inserta nuevo espacio  en fila__________________________________
func ajustaraFila(fila *NodoMatris, aInsertar *NodoMatris) {
	nodoActual := fila
	for nodoActual.getPosterior() != nil && nodoActual.PosX < aInsertar.PosX {
		nodoActual = nodoActual.getPosterior()
	}
	if nodoActual.PosX < aInsertar.PosX {
		if nodoActual.getPosterior() != nil {
			n2 := nodoActual.getPosterior()
			nodoActual.setPosterior(aInsertar)
			n2.setAnterior(aInsertar)
			aInsertar.setAnterior(nodoActual)
			aInsertar.setPosterior(n2)
		} else {
			nodoActual.setPosterior(aInsertar)
			aInsertar.setAnterior(nodoActual)
		}
	}
}

// _____________________________inserta nuevo espacio  columna__________________________________
func ajustaraColumna(columna *NodoMatris, aInsertar *NodoMatris) {
	nodoActual := columna
	for nodoActual.getInferior() != nil && nodoActual.PosY < aInsertar.PosY {
		nodoActual = nodoActual.getInferior()
	}
	if nodoActual.PosY < aInsertar.PosY {
		if nodoActual.getInferior() != nil {
			n2 := nodoActual.getInferior()
			nodoActual.setInferior(aInsertar)
			n2.setSuperior(aInsertar)
			aInsertar.setSuperior(nodoActual)
			aInsertar.setInferior(n2)
		} else {
			nodoActual.setInferior(aInsertar)
			aInsertar.setSuperior(nodoActual)
		}
	}
}

// ______________________________________graficar__________________________________________

func (m *MatrisDispersa) declararyUnirCabezeras() string {
	texto := ""
	//se declaran los nodos de la cabexzera de columna
	actualColumna := m.CabezeraColumnas
	for actualColumna != nil {
		texto += "nodo" + actualColumna.getId() + " [label=\"" + actualColumna.Dia + "\"] \n"
		actualColumna = actualColumna.getPosterior()
	}
	//se declaran los nodos de la cabezera de filas
	actualFila := m.CabezeraFilas
	for actualFila != nil {
		texto += "nodo" + actualFila.getId() + " [label=\"" + actualFila.Departamento + "\"] \n"
		actualFila = actualFila.getInferior()
	}
	//se unen los nodos de la cabezera de columnas
	actualColumna = m.CabezeraColumnas
	texto += "Matris -> nodo" + actualColumna.getId() + " [dir=both] \n"
	for actualColumna.getPosterior() != nil {
		texto += "nodo" + actualColumna.getId() + " -> nodo" + actualColumna.posterior.getId() + " [dir=both] \n"
		actualColumna = actualColumna.getPosterior()
	}
	//se unen los nodos de la cabezera de filas
	actualFila = m.CabezeraFilas
	texto += "Matris -> nodo" + actualFila.getId() + " [dir=both] \n"
	for actualFila.getInferior() != nil {
		texto += "nodo" + actualFila.getId() + " -> nodo" + actualFila.inferior.getId() + " [dir=both] \n"
		actualFila = actualFila.getInferior()
	}
	return texto
}

// ______________________________________declarar nodos Internos__________________________________________

func (m *MatrisDispersa) declararNodosInternos() string {
	texto := ""
	columnaActual := m.CabezeraColumnas
	for columnaActual != nil {
		texto += getTextoNodosColumna(columnaActual.getInferior())
		columnaActual = columnaActual.getPosterior()
	}
	return texto
}
func getTextoNodosColumna(nodo *NodoMatris) string {
	if nodo == nil {
		return "\n"
	}
	if nodo.getInferior() != nil {
		return "nodo" + nodo.getId() + "[label=\"" + strconv.Itoa(nodo.cola.Longitud) + "\"] \n" + getTextoNodosColumna(nodo.getInferior())
	}
	return "nodo" + nodo.getId() + "[label=\"" + strconv.Itoa(nodo.cola.Longitud) + "\"] \n"
}

// ___________________________________ unir nodos verticalmente_______________________________
func (m *MatrisDispersa) unirNodosVertical() string {
	texto := ""
	columnaActual := m.CabezeraColumnas
	for columnaActual != nil {
		texto += unirVertical(columnaActual)
		columnaActual = columnaActual.getPosterior()
	}
	return texto
}
func unirVertical(nodo *NodoMatris) string {
	if nodo.getInferior() != nil {
		return "nodo" + nodo.getId() + " -> nodo" + nodo.getInferior().getId() + "[dir=both] \n" + unirVertical(nodo.getInferior())
	}
	return ""
}

// ___________________________________ unir nodos horizontalmente_______________________________
func (m *MatrisDispersa) unirNodosHorizontal() string {
	texto := ""
	filaActual := m.CabezeraFilas
	for filaActual != nil {
		texto += unirHorizontal(filaActual)
		filaActual = filaActual.getInferior()
	}
	return texto
}
func unirHorizontal(nodo *NodoMatris) string {
	if nodo == nil {
		return "\n"
	}
	if nodo.getPosterior() != nil {
		return "nodo" + nodo.getId() + " -> nodo" + nodo.getPosterior().getId() + "[constraint=false, dir=both] \n" + unirHorizontal(nodo.getPosterior())
	}
	return ""
}

// ___________________________________ establecer nivel de filas_______________________________
func (m *MatrisDispersa) establecerNivelFilas() string {
	texto := "{ rank=same; Matris"
	columnaActual := m.CabezeraColumnas
	for columnaActual != nil {
		texto += "; nodo" + columnaActual.getId()
		columnaActual = columnaActual.getPosterior()
	}
	texto += "} \n"

	filaActual := m.CabezeraFilas
	for filaActual != nil {
		texto += "{ rank=same;" + nivelFilas(filaActual) + "} \n"
		filaActual = filaActual.getInferior()
	}
	return texto
}
func nivelFilas(nodo *NodoMatris) string {
	if nodo == nil {
		return ""
	}
	if nodo.getPosterior() != nil {
		return "nodo" + nodo.getId() + ";" + nivelFilas(nodo.getPosterior())
	}
	return "nodo" + nodo.getId()

}

// _________________________________ Insertar __________________________________

//Insertar ...
func (m *MatrisDispersa) Insertar(pedido pedidos.Pedido, Departamento string, Dia string) {
	if m.CabezeraColumnas == nil {
		m.CabezeraColumnas = &NodoMatris{EsCabezera: true, Dia: Dia, PosY: -1}
	}
	if m.CabezeraFilas == nil {
		m.CabezeraFilas = &NodoMatris{EsCabezera: true, Departamento: Departamento, PosX: -1}
	}
	posibleNuevaColumna := &NodoMatris{EsCabezera: true, Dia: Dia, PosY: -1}
	posibleNuevaFila := &NodoMatris{EsCabezera: true, Departamento: Departamento, PosX: -1}
	columna := m.getEncabezadoVertical(Dia, posibleNuevaColumna)
	fila := m.getEncabezadoHorizontal(Departamento, posibleNuevaFila)
	if actualizar(columna, fila, pedido) == false {
		c := new(Cola)
		c.Encolar(pedido)
		nodo := &NodoMatris{cola: *c, PosX: columna.PosX, PosY: fila.PosY}
		ajustaraColumna(columna, nodo)
		ajustaraFila(fila, nodo)
	}
}

// _________________________________ Graficar __________________________________

//Graficar
func (m *MatrisDispersa) Graficar(nombre string) {
	if m.CabezeraColumnas == nil || m.CabezeraFilas == nil {
		return
	}
	texto := "digraph { \n"
	texto += "rankdir = TB; \n"
	texto += "node [shape=rectangle, height=0.5, width=2] \n"
	texto += "graph[ nodesep = 0.5] \n"
	texto += "Matris[label=\"Matris\"] \n"
	texto += m.declararyUnirCabezeras()
	texto += m.declararNodosInternos()
	texto += m.unirNodosVertical()
	texto += m.unirNodosHorizontal()
	texto += m.establecerNivelFilas()
	texto += "}"
	_ = ioutil.WriteFile(nombre+".dot", []byte(texto), 0644)
	cmd := exec.Command("dot", "-Tjpg", nombre+".dot", "-o", "./pedidos/"+nombre+".jpg")
	_ = cmd.Run()
	_ = os.Remove(nombre + ".dot")
}

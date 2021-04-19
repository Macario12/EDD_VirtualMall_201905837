package arbol_avl

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type Inventario struct {
	Nombre         string `json:"Nombre"`
	Codigo         int    `json:"Codigo"`
	Descripcion    string `json:"Descripcion"`
	Precio         int    `json:"Precio"`
	Cantidad       int    `json:"Cantidad"`
	Imagen         string `json:"Imagen"`
	Almacenamiento string `json:"Almacenamiento"`
}

func NewInventario(nombre string, codigo int, descrip string, precio int, cantidad int, imagen string, almacenamiento string) *Inventario {
	return &Inventario{
		Nombre:         nombre,
		Codigo:         codigo,
		Descripcion:    descrip,
		Precio:         precio,
		Cantidad:       cantidad,
		Imagen:         imagen,
		Almacenamiento: almacenamiento,
	}
}

type Node struct {
	Data   Inventario
	Height int
	Left   *Node
	Right  *Node
}

type Tree struct {
	Root *Node
}

func Newtree() *Tree {
	return &Tree{nil}
}

func Newnode(data Inventario) *Node {
	return &Node{data, 0, nil, nil}
}

func Rotateleft(n *Node, n1 *Node) *Node {
	n.Left = n1.Right
	n1.Right = n

	if n1.Height == -1 {
		n.Height = 0
		n1.Height = 0
	} else {
		n.Height = -1
		n1.Height = 1
	}
	return n1

}

func Rotateright(n *Node, n1 *Node) *Node {

	n.Right = n1.Left
	n1.Left = n

	if n1.Height == 1 {
		n.Height = 0
		n1.Height = 0
	} else {
		n.Height = 1
		n1.Height = -1
	}

	return n1
}

func Rotaterightleft(n *Node, n1 *Node) *Node {
	n2 := n1.Left
	n.Right = n2.Left
	n2.Left = n
	n1.Left = n2.Right
	n2.Right = n1

	if n2.Height == 1 {
		n.Height = -1
	} else {
		n.Height = 0
	}

	if n2.Height == -1 {
		n1.Height = 1

	} else {
		n1.Height = 0
	}

	n2.Height = 0

	return n2
}

func Rotateleftright(n *Node, n1 *Node) *Node {
	n2 := n1.Right
	n.Left = n2.Right
	n2.Right = n

	n1.Right = n2.Left
	n2.Left = n1

	if n2.Height == 1 {
		n1.Height = -1
	} else {
		n1.Height = 0
	}

	if n2.Height == -1 {
		n.Height = 1
	} else {
		n.Height = 0
	}

	n2.Height = 0

	return n2

}

func Insert(rootr *Node, dataa Inventario, hc *bool) *Node {
	var n1 *Node

	if rootr == nil {
		rootr = Newnode(dataa)
		*hc = true
	} else if dataa.Codigo < rootr.Data.Codigo {
		lef := Insert(rootr.Left, dataa, hc)
		rootr.Left = lef
		if *hc {
			switch rootr.Height {
			case 1:
				rootr.Height = 0
				*hc = false
				break
			case 0:
				rootr.Height = -1
				break
			case -1:
				n1 = rootr.Left

				if n1.Height == -1 {
					rootr = Rotateleft(rootr, n1)
				} else {
					rootr = Rotateleftright(rootr, n1)
				}
				*hc = false
			}
		}
	} else if dataa.Codigo > rootr.Data.Codigo {
		der := Insert(rootr.Right, dataa, hc)
		rootr.Right = der

		if *hc {
			switch rootr.Height {
			case 1:
				n1 = rootr.Right
				if n1.Height == 1 {
					rootr = Rotateright(rootr, n1)

				} else {
					rootr = Rotaterightleft(rootr, n1)
				}
				*hc = false
				break
			case 0:
				rootr.Height = 1
				break
			case -1:
				rootr.Height = 0
				*hc = false
			}
		}
	} else if dataa.Codigo == rootr.Data.Codigo {
		rootr.Data.Cantidad += dataa.Cantidad
	}

	return rootr
}

func (this *Tree) Insertroot(dataa Inventario) {
	b := false
	a := &b

	this.Root = Insert(this.Root, dataa, a)
}

func PreOrden(n *Node) {
	if n != nil {
		fmt.Println(n.Data.Codigo, " ")
		PreOrden(n.Left)
		PreOrden(n.Right)
	}
}

func Graficainterna(arboli *Node) string {
	var etiqueta string
	if arboli.Left == nil && arboli.Right == nil {

		etiqueta = "nodo" + strconv.Itoa(arboli.Data.Codigo) + "[ shape=record, label =\"{" + "Codigo: " + strconv.Itoa(arboli.Data.Codigo) + "|" + "Nombre: " + arboli.Data.Nombre + "|" + "Precio: " + strconv.Itoa(arboli.Data.Precio) + "|" + "Cantidad: " + strconv.Itoa(arboli.Data.Cantidad) + "}\"];\n"
	} else {

		etiqueta = "nodo" + strconv.Itoa(arboli.Data.Codigo) + "[ shape=record, label =\"{" + "Codigo: " + strconv.Itoa(arboli.Data.Codigo) + "|" + "Nombre: " + arboli.Data.Nombre + "|" + "Precio: " + strconv.Itoa(arboli.Data.Precio) + "|" + "Cantidad: " + strconv.Itoa(arboli.Data.Cantidad) + "}\"];\n"
	}

	if arboli.Left != nil {
		etiqueta = etiqueta + Graficainterna(arboli.Left) + "nodo" + strconv.Itoa(arboli.Data.Codigo) + "->nodo" + strconv.Itoa(arboli.Left.Data.Codigo) + "\n"
	}
	if arboli.Right != nil {
		etiqueta = etiqueta + Graficainterna(arboli.Right) + "nodo" + strconv.Itoa(arboli.Data.Codigo) + "->nodo" + strconv.Itoa(arboli.Right.Data.Codigo) + "\n"
	}

	return etiqueta
}

func Graficararbol(nodoraiz *Node, nombre string) string {
	archivo, _ := os.Create("./graficoArboles/" + nombre + ".dot")
	_, _ = archivo.WriteString("digraph grafico{" + "\n")
	_, _ = archivo.WriteString("rankdir=UD \n")
	_, _ = archivo.WriteString("node[shape=box] \n")
	_, _ = archivo.WriteString("concentrate=true \n")
	_, _ = archivo.WriteString(Graficainterna(nodoraiz))
	_, _ = archivo.WriteString("}")
	archivo.Close()
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./graficoArboles/"+nombre+".dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("./graficoArboles/"+nombre+".png", cmd, os.FileMode(mode))

	f, _ := os.Open("./graficoArboles/" + nombre + ".png")

	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded
}

var Datobuscado Inventario

func busquedaProducto(n *Node, datobuscar Inventario) {
	if n != nil {

		if datobuscar.Codigo == n.Data.Codigo {
			if datobuscar.Cantidad <= n.Data.Cantidad {
				n.Data.Cantidad = n.Data.Cantidad - datobuscar.Cantidad
				Datobuscado = n.Data
			} else {
				n.Data.Cantidad = 0
				Datobuscado = n.Data
			}
		}

		busquedaProducto(n.Left, datobuscar)
		busquedaProducto(n.Right, datobuscar)
	}
}

/*
func main() {
	a := Newtree()
	a.Insertroot(*NewInventario("A", 4511, "", 4, 5, ""))
	a.Insertroot(*NewInventario("B", 4512, "", 4, 1, ""))
	a.Insertroot(*NewInventario("C", 4513, "", 4, 2, ""))
	a.Insertroot(*NewInventario("D", 4514, "", 4, 3, ""))
	a.Insertroot(*NewInventario("E", 4515, "", 4, 5, ""))
	a.Insertroot(*NewInventario("F", 4516, "", 4, 6, ""))
	a.Insertroot(*NewInventario("D", 4514, "", 4, 7, ""))
	a.Insertroot(*NewInventario("H", 4517, "", 4, 9, ""))
	a.Insertroot(*NewInventario("I", 4518, "", 4, 2, ""))
	a.Insertroot(*NewInventario("J", 4519, "", 4, 56, ""))
	a.Insertroot(*NewInventario("B", 4512, "", 4, 51, ""))

	fmt.Println("eliminado")

	PreOrden(a.Root)

}*/

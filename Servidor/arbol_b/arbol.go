package arbol_b

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Arbol struct {
	k    int
	Raiz *Nodo
}

func NewArbol(nivel int) *Arbol {
	a := Arbol{nivel, nil}
	nodoRaiz := NewNodo(nivel)
	a.Raiz = nodoRaiz

	return &a
}
type Lave struct{
 Llave string `json:"LLave"`
}

var Llave Lave

func (this *Arbol) Insertar(newKey *Key) {
	if this.Raiz.Keys[0] == nil {
		this.Raiz.Colocar(0, newKey)
	} else if this.Raiz.Keys[0].Izquierdo == nil {
		lugarinsertado := -1
		node := this.Raiz

		lugarinsertado = this.colocarNodo(node, newKey)

		if lugarinsertado != -1 {
			if lugarinsertado == node.Max-1 {
				mid := node.Max / 2
				llavecentral := node.Keys[mid]

				derecho := NewNodo(this.k)

				izquierdo := NewNodo(this.k)

				indiceizq := 0
				indicedere := 0

				for j := 0; j < node.Max; j++ {
					if node.Keys[j].Value.DPI < llavecentral.Value.DPI {
						izquierdo.Colocar(indiceizq, node.Keys[j])
						indiceizq++
						node.Colocar(j, nil)
					} else if node.Keys[j].Value.DPI > llavecentral.Value.DPI {
						derecho.Colocar(indicedere, node.Keys[j])
						indicedere++
						node.Colocar(j, nil)

					}

				}

				node.Colocar(mid, nil)
				this.Raiz = node
				this.Raiz.Colocar(0, llavecentral)
				izquierdo.NodoPadre = this.Raiz
				derecho.NodoPadre = this.Raiz

				llavecentral.Izquierdo = izquierdo
				llavecentral.Derecho = derecho
			}
		}

	} else if this.Raiz.Keys[0].Izquierdo != nil {
		node := this.Raiz

		for node.Keys[0].Izquierdo != nil {
			loop := 0

			for i := 0; i < node.Max; i, loop = i+1, loop+1 {
				if node.Keys[i] != nil {
					if node.Keys[i].Value.DPI > newKey.Value.DPI {
						node = node.Keys[i].Izquierdo
						break

					}
				} else {
					node = node.Keys[i-1].Derecho
					break
				}
			}

			if loop == node.Max {
				node = node.Keys[loop-1].Derecho
			}
		}

		indiceColocado := this.colocarNodo(node, newKey)

		if indiceColocado == node.Max-1 {
			for node.NodoPadre != nil {
				indicemedio := node.Max / 2

				llavecentral := node.Keys[indicemedio]

				izquierdo := NewNodo(this.k)

				derecho := NewNodo(this.k)

				indiceizquierdo, indicederecho := 0, 0

				for i := 0; i < node.Max; i++ {
					if node.Keys[i].Value.DPI < llavecentral.Value.DPI {
						izquierdo.Colocar(indiceizquierdo, node.Keys[i])
						indiceizquierdo++
						node.Colocar(i, nil)
					} else if node.Keys[i].Value.DPI > llavecentral.Value.DPI {
						derecho.Colocar(indicederecho, node.Keys[i])
						indicederecho++

						node.Colocar(i, nil)
					}
				}

				node.Colocar(indicemedio, nil)
				llavecentral.Izquierdo = izquierdo

				llavecentral.Derecho = derecho

				node = node.NodoPadre

				izquierdo.NodoPadre = node

				derecho.NodoPadre = node

				for i := 0; i < izquierdo.Max; i++ {
					if izquierdo.Keys[i] != nil {
						if izquierdo.Keys[i].Izquierdo != nil {
							izquierdo.Keys[i].Izquierdo.NodoPadre = izquierdo
						}

						if izquierdo.Keys[i].Derecho != nil {
							izquierdo.Keys[i].Derecho.NodoPadre = izquierdo
						}
					}

				}

				for i := 0; i < derecho.Max; i++ {

					if derecho.Keys[i] != nil {
						if derecho.Keys[i].Izquierdo != nil {
							derecho.Keys[i].Izquierdo.NodoPadre = derecho
						}

						if derecho.Keys[i].Derecho != nil {
							derecho.Keys[i].Derecho.NodoPadre = derecho
						}
					}

				}

				lugarcolocado := this.colocarNodo(node, llavecentral)

				if lugarcolocado == node.Max-1 {
					if node.NodoPadre == nil {
						indicecentralraiz := node.Max / 2

						llavecentralraiz := node.Keys[indicecentralraiz]
						izquierdoraiz := NewNodo(this.k)
						derechoraiz := NewNodo(this.k)
						indicederechoraiz, indiceizquierdoraiz := 0, 0

						for i := 0; i < node.Max; i++ {
							if node.Keys[i].Value.DPI < llavecentralraiz.Value.DPI {
								izquierdoraiz.Colocar(indiceizquierdoraiz, node.Keys[i])

								indiceizquierdoraiz++
								node.Colocar(i, nil)

							} else if node.Keys[i].Value.DPI > llavecentralraiz.Value.DPI {
								derechoraiz.Colocar(indicederechoraiz, node.Keys[i])
								indicederechoraiz++
								node.Colocar(i, nil)
							}

						}

						node.Colocar(indicecentralraiz, nil)
						node.Colocar(0, llavecentralraiz)

						for i := 0; i < this.k; i++ {
							if izquierdoraiz.Keys[i] != nil {
								izquierdoraiz.Keys[i].Izquierdo.NodoPadre = izquierdoraiz
								izquierdoraiz.Keys[i].Derecho.NodoPadre = izquierdoraiz
							}

						}

						for i := 0; i < this.k; i++ {
							if derechoraiz.Keys[i] != nil {
								derechoraiz.Keys[i].Izquierdo.NodoPadre = derechoraiz
								derechoraiz.Keys[i].Derecho.NodoPadre = derechoraiz
							}
						}

						llavecentralraiz.Izquierdo = izquierdoraiz

						llavecentralraiz.Derecho = derechoraiz

						izquierdoraiz.NodoPadre = node
						derechoraiz.NodoPadre = node

						this.Raiz = node
					}

					continue
				} else {
					break
				}
			}
		}
	}
}

func (this *Arbol) colocarNodo(nodo *Nodo, newKey *Key) int {
	index := -1
	for i := 0; i < nodo.Max; i++ {
		if nodo.Keys[i] == nil {
			placed := false
			for j := i - 1; j >= 0; j-- {
				if nodo.Keys[j].Value.DPI > newKey.Value.DPI {
					nodo.Colocar(j+1, nodo.Keys[j])
				} else {
					nodo.Colocar(j+1, newKey)
					nodo.Keys[j].Derecho = newKey.Izquierdo
					if (j+2) < this.k && nodo.Keys[j+2] != nil {
						nodo.Keys[j+2].Izquierdo = newKey.Derecho
					}

					placed = true
					break
				}
			}

			if placed == false {
				nodo.Colocar(0, newKey)
				nodo.Keys[1].Izquierdo = newKey.Derecho

			}

			index = i
			break
		}
	}
	return index
}

func (this *Arbol) GraficarArbolSin() string {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "digraph G{\nnode[shape=record]\nedge[color=\"greenyellow\"]\n")
	m := make(map[string]*Nodo)
	graficar(this.Raiz, &builder, m, nil, 0)
	fmt.Fprintf(&builder, "}")
	guardarArchivo(builder.String())

	return generarImagen("arbol.png")
}

func (this *Arbol) GraficarArbolCifrado() string {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "digraph G{\nnode[shape=record]\nedge[color=\"greenyellow\"]\n")
	m := make(map[string]*Nodo)
	graficarCifrado(this.Raiz, &builder, m, nil, 0)
	fmt.Fprintf(&builder, "}")
	guardarArchivo(builder.String())

	return generarImagen("arbolCifrado.png")
}

func (this *Arbol) GraficarArbolCifradoSencible() string {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "digraph G{\nnode[shape=record]\nedge[color=\"greenyellow\"]\n")
	m := make(map[string]*Nodo)
	graficarCifradoSencible(this.Raiz, &builder, m, nil, 0)
	fmt.Fprintf(&builder, "}")
	guardarArchivo(builder.String())

	return generarImagen("arbolCifradoSencible.png")
}

func graficar(actual *Nodo, cad *strings.Builder, arr map[string]*Nodo, padre *Nodo, pos int) {
	if actual == nil {
		return
	}
	j := 0
	contiene := arr[fmt.Sprint(&(*actual))]
	if contiene != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad, "node%p[color=\".7 .3 1.0\",label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++

			} else {
				fmt.Fprintf(cad, "{<f%d>DPI: %d|", j, actual.Keys[i].Value.DPI)
				fmt.Fprintf(cad, "Nombre: %s|", actual.Keys[i].Value.Nombre)
				fmt.Fprintf(cad, "Correo: %s|", actual.Keys[i].Value.Correo)
				fmt.Fprintf(cad, "Password: %s|", actual.Keys[i].Value.Password)
				fmt.Fprintf(cad, "Cuenta: %s}|", actual.Keys[i].Value.Cuenta)
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad, "\"]\n")
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		graficar(actual.Keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficar(actual.Keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p:f%d->node%p\n", &(*padre), pos, &(*actual))
	}
}
func graficarCifrado(actual *Nodo, cad *strings.Builder, arr map[string]*Nodo, padre *Nodo, pos int) {
	if actual == nil {
		return
	}
	j := 0
	contiene := arr[fmt.Sprint(&(*actual))]
	if contiene != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad, "node%p[color=\".7 .3 1.0\",label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++

			} else {
				fmt.Fprintf(cad, "{<f%d>DPI: %s|", j, encrypt(strconv.Itoa(actual.Keys[i].Value.DPI), Llave.Llave))
				fmt.Fprintf(cad, "Nombre: %s|", encrypt(actual.Keys[i].Value.Nombre,Llave.Llave))
				fmt.Fprintf(cad, "Correo: %s|", encrypt(actual.Keys[i].Value.Correo,Llave.Llave))
				fmt.Fprintf(cad, "Password: %s|", actual.Keys[i].Value.Password)
				fmt.Fprintf(cad, "Cuenta: %s}|", encrypt(actual.Keys[i].Value.Cuenta, Llave.Llave))
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad, "\"]\n")
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		graficarCifrado(actual.Keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficarCifrado(actual.Keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p:f%d->node%p\n", &(*padre), pos, &(*actual))
	}
}
func graficarCifradoSencible(actual *Nodo, cad *strings.Builder, arr map[string]*Nodo, padre *Nodo, pos int) {
	if actual == nil {
		return
	}
	j := 0
	contiene := arr[fmt.Sprint(&(*actual))]
	if contiene != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad, "node%p[color=\".7 .3 1.0\",label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++

			} else {
				fmt.Fprintf(cad, "{<f%d>DPI: %s|", j, encrypt(strconv.Itoa(actual.Keys[i].Value.DPI), Llave.Llave))
				fmt.Fprintf(cad, "Nombre: %s|", actual.Keys[i].Value.Nombre)
				fmt.Fprintf(cad, "Correo: %s|", encrypt(actual.Keys[i].Value.Correo,Llave.Llave))
				fmt.Fprintf(cad, "Password: %s|", actual.Keys[i].Value.Password)
				fmt.Fprintf(cad, "Cuenta: %s}|", actual.Keys[i].Value.Cuenta)
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad, "\"]\n")
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		graficarCifradoSencible(actual.Keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficarCifradoSencible(actual.Keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p:f%d->node%p\n", &(*padre), pos, &(*actual))
	}
}
func guardarArchivo(cadena string) {
	f, err := os.Create("diagrama.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(cadena)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written succesfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func generarImagen(nombre string) string {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./diagrama.dot").Output()
	mode := int(0772)
	ioutil.WriteFile(nombre, cmd, os.FileMode(mode))
	f, _ := os.Open("./" + nombre)

	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded
}

var usuario User

func Buscar(actual *Nodo, dato User, padre *Nodo, pos int) User {

	if actual == nil {
		return usuario
	}
	j := 0
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return usuario
		} else {
			if enlace {

				enlace = false
				i--
				j++

			} else {

				if dato.DPI == actual.Keys[i].Value.DPI || dato.Correo == actual.Keys[i].Value.Correo {
					if dato.Password == actual.Keys[i].Value.Password {
						usuario = actual.Keys[i].Value
					} else {
					}
				} else {
				}
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {

						j++
						break
					}
				}
			}
		}
	}
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		Buscar(actual.Keys[i].Izquierdo, dato, actual, ji)
		ji++
		ji++
		Buscar(actual.Keys[i].Derecho, dato, actual, ji)
		ji++
		ji--
	}
	if padre != nil {

	}

	return usuario

}

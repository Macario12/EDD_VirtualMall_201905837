package Tabla_Hash

import (
	"fmt"
	"time"

	"../arbol_b"
)

// ____________TABLA HASH _____________

type TablaHash struct {
	Id         string
	IdentLevel int
	Registros  []*Registro
	Capacidad  int
	Carga      int
}

func NewTablaHash(cap int) *TablaHash {
	return &TablaHash{
		Id:        time.Now().String(),
		Registros: make([]*Registro, cap, cap),
		Capacidad: cap,
	}
}

// ___________ REGISTRO ____________
type Registro struct {
	IdContenedor string
	Usuario      *arbol_b.User
	IdentLevel   int
	Comentario   string
	TablaHash    *TablaHash
}

func (t *TablaHash) NewRegistro(usuario  *arbol_b.User, comentario string) *Registro {
	fmt.Println("Entras")
	tablaHija := NewTablaHash(t.Capacidad)
	tablaHija.IdentLevel = t.IdentLevel + 1
	return &Registro{t.Id, usuario, t.IdentLevel, comentario, tablaHija}
}

// ___________ FUNCIONALIDADES ____________
func isPrime(num int) bool {
	contDivisores := 0
	for i := 1; i <= int(num/2); i++ {
		if num%i == 0 {
			contDivisores++
		}
		if contDivisores > 1 {
			return false
		}
	}
	return true
}

func (t *TablaHash) FuncHashMult(clave int) int {
	return t.Capacidad * ((clave * (3 / 7)) % 1)
}

func (t *TablaHash) ExplorCuadratic(clave, intento int) int {
	return (t.Capacidad*((clave*(3/7))%1) + (intento * intento)) % t.Capacidad
}

func (t *TablaHash) PosOcupada(pos int) bool {
	if t.Registros[pos] != nil {
		return true
	}
	return false
}

func (t *TablaHash) getFactoCarga() float64 {
	return float64(t.Carga) / float64(t.Capacidad)
}

func (t *TablaHash) resize() {
	if t.getFactoCarga() < 0.5 {
		return
	}
	newSize := t.Capacidad + 1
	for !isPrime(newSize) {
		newSize++
	}
	for i := 0; i < newSize-t.Capacidad; i++ {
		t.Registros = append(t.Registros, nil)
	}
	t.Capacidad = newSize
	//fmt.Println("Hubo resizing a un tamaño de: " + strconv.Itoa(len(t.Registros)))
}

func (t *TablaHash) Insertar(clave int, reg *Registro) {
	//fmt.Println("Entro")
	pos := t.FuncHashMult(clave)
	intento := 1
	for t.PosOcupada(pos) {
		//fmt.Println("Hubo colision con la posicion: " + strconv.Itoa(pos))
		pos = t.ExplorCuadratic(clave, intento)
		intento++
	}
	//fmt.Println(pos)
	reg.IdentLevel = t.IdentLevel
	t.Registros[pos] = reg
	t.Carga++
	t.resize()
}

func (t *TablaHash) GetTabla(idTabla string) *TablaHash {
	if t.Id == idTabla {
		return t
	}
	for _, Registro := range t.Registros {
		if Registro != nil {
			tabla := Registro.TablaHash.GetTabla(idTabla)
			if tabla != nil {
				return tabla
			}
		}
	}
	return nil
}

func (t *TablaHash) GetAsList() []*Registro {
	var listaRegistros []*Registro
	for _, registro := range t.Registros {
		if registro != nil {
			listaRegistros = append(listaRegistros, registro)
			listaRegistros = append(listaRegistros, registro.TablaHash.GetAsList()...)
		}
	}
	return listaRegistros
}

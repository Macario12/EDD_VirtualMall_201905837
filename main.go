package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"strconv"

	"os"

	"./estructura"
	"./tienda"
	"github.com/gorilla/mux"
)

type vector struct {
	Datos []datos `json:"Datos"`
}

type datos struct {
	Indice        string          `json:"Indice"`
	Departamentos []departamentos `json:"Departamentos"`
}

type departamentos struct {
	Departamento string  `json:"Nombre"`
	Tiendas      []Store `json:"Tiendas"`
}

type Store struct {
	Name        string `json:"Nombre"`
	Description string `json:"Descripcion"`
	Contact     string `json:"Contacto"`
	Score       int    `json:"Calificacion"`
}

type structaux struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}

var vectorJSON vector
var linealizar []estructura.List

func agregar(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &vectorJSON)

	json.NewEncoder(w).Encode(vectorJSON)
	linealizarMatriz()
}

var indicesfor []string
var depa []string

func linealizarMatriz() {
	linealizar = make([]estructura.List, 0, len(vectorJSON.Datos)*len(vectorJSON.Datos[0].Departamentos)*5)
	var varpureba int
	for i := 0; i < len(vectorJSON.Datos); i++ {
		fmt.Println("Indice: " + (vectorJSON.Datos[i].Indice))
		indicesfor = append(indicesfor, (vectorJSON.Datos[i].Indice))
		for j := 0; j < len(vectorJSON.Datos[i].Departamentos); j++ {
			fmt.Println("Departamento: " + vectorJSON.Datos[i].Departamentos[j].Departamento)
			if varpureba == 0 {
				depa = append(depa, vectorJSON.Datos[i].Departamentos[j].Departamento)
			}

			var lista1 estructura.List
			var lista2 estructura.List
			var lista3 estructura.List
			var lista4 estructura.List
			var lista5 estructura.List
			for x := 0; x < len(vectorJSON.Datos[i].Departamentos[j].Tiendas); x++ {
				nombreaux := vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Name
				descraux := vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Description
				cotactaux := vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Contact
				scoreaux := vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Score

				tiendaaux := tienda.Store{nombreaux, descraux, cotactaux, scoreaux}
				if scoreaux == 1 {
					lista1.Add(tiendaaux)

				} else if scoreaux == 2 {
					lista2.Add(tiendaaux)

				} else if scoreaux == 3 {
					lista3.Add(tiendaaux)
				} else if scoreaux == 4 {
					lista4.Add(tiendaaux)
				} else if scoreaux == 5 {
					lista5.Add(tiendaaux)
				}

			}
			linealizar = append(linealizar, lista1)
			linealizar = append(linealizar, lista2)
			linealizar = append(linealizar, lista3)
			linealizar = append(linealizar, lista4)
			linealizar = append(linealizar, lista5)
		}

		varpureba++
	}

	fmt.Println(depa)

	for i := 0; i < len(linealizar); i++ {
		Ordenar(linealizar[i])
	}

	for i := 0; i < len(linealizar); i++ {
		imprimir(linealizar[i])
	}
}

func Ordenar(l estructura.List) {
	aux := l.Frist
	temp := tienda.Store{}

	for aux != nil {
		aux2 := aux.GetNext()
		for aux2 != nil {
			if codigoASCII(aux2.Store.Name) > codigoASCII(aux.Store.Name) {
				temp = aux.Store
				aux.Store = aux2.Store
				aux2.Store = temp
			}
			aux2 = aux2.GetNext()
		}
		aux = aux.GetNext()

	}
}

func busquedaPosicionLinealizada(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	posicion, _ := strconv.Atoi(vars["id"])
	listaObtenida := linealizar[posicion]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)

	a := listaObtenida.Frist
	var sliceTiendas []tienda.Store

	for a != nil {
		tienda := tienda.Store{a.Store.Name, a.Store.Description, a.Store.Contact, a.Store.Score}
		sliceTiendas = append(sliceTiendas, tienda)
		a = a.Next
	}

	json.NewEncoder(w).Encode(sliceTiendas)
}

func imprimir(l estructura.List) {
	a := l.Frist

	for a != nil {
		fmt.Println(a.Store.Name)
		fmt.Println(a.Store.Description)
		fmt.Println(a.Store.Contact)
		fmt.Println(strconv.Itoa(a.Store.Score))

		a = a.Next
	}
}

func codigoASCII(nombre string) int {
	cadena := []rune(nombre)

	var ascii int

	for i := 0; i < len(cadena); i++ {
		ascii += int(cadena[i])
	}
	return ascii
}

func busquedaEspecificaTienda(w http.ResponseWriter, r *http.Request) {
	var varaux structaux

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error al enviar")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &varaux)
	indice := string(varaux.Nombre[0])
	var posicionL int
	var posi, posj int
	for i := 0; i < len(indicesfor); i++ {
		if indice == (indicesfor[i]) {
			posi = i
		}
	}
	for j := 0; j < len(depa); j++ {
		if varaux.Departamento == (depa[j]) {
			posj = j
		}
	}
	posicionL = ((posi*len(depa))+posj)*5 + (varaux.Calificacion - 1)
	json.NewEncoder(w).Encode(linealizar[posicionL].SearchStore(varaux.Nombre))
}

func eliminarTienda(w http.ResponseWriter, r *http.Request) {
	var varaux structaux

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error al enviar")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &varaux)
	indice := string(varaux.Nombre[0])
	var posicionL int
	var posi, posj int
	for i := 0; i < len(indicesfor); i++ {
		if indice == (indicesfor[i]) {
			posi = i
		}
	}
	for j := 0; j < len(depa); j++ {
		if varaux.Departamento == (depa[j]) {
			posj = j
		}
	}
	posicionL = ((posi*len(depa))+posj)*5 + (varaux.Calificacion - 1)
	if linealizar[posicionL].DeleteStore(varaux.Nombre) {

		json.NewEncoder(w).Encode("Se eliminó exitosamente")
	} else {
		json.NewEncoder(w).Encode("que chingue a su madre el america")
	}
}

func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Funcionan EDD")
}

func graficar(w http.ResponseWriter, r *http.Request) {
	var contadorGraficapart int
	/*	for i := 0; i < len(vectorJSON.Datos); i++ {
		for j := 0; j < len(vectorJSON.Datos[i].Departamentos); j++ {
			contadorGraficapart++
			graficodelArreglo(i, j, contadorGraficapart)
		}
	}*/

	for i := 0; i < len(indicesfor); i++ {
		for j := 0; j < len(depa); j++ {
			contadorGraficapart++
			graficodelArreglo(i, j, contadorGraficapart)
		}
	}

	fmt.Fprintf(w, "to cool")
	iniciociclo = 0
	posicionesvector = 0
}

var iniciociclo int
var posicionesvector int

func graficodelArreglo(i int, j int, contadorGraficapart int) {
	archivo, _ := os.Create("graficoLinealizado" + strconv.Itoa(contadorGraficapart) + ".dot")
	_, _ = archivo.WriteString("digraph grafico{" + "\n")
	_, _ = archivo.WriteString("compound=true;" + "\n")
	_, _ = archivo.WriteString("subgraph cluster0{" + "\n")
	_, _ = archivo.WriteString("edge[minlen=0.1, dir=fordware]" + "\n")
	var contador int
	var contadoraux int
	for x := 0; x < 5; x++ {

		_, _ = archivo.WriteString("struct" + strconv.Itoa(contador) + "[shape=record,label=\"" + (vectorJSON.Datos[i].Indice) + "|" + vectorJSON.Datos[i].Departamentos[j].Departamento + "|{" + strconv.Itoa(x+1) + "| pos:" + strconv.Itoa(posicionesvector) + "}\"];" + "\n")
		posicionesvector++
		contador++

	}

	for x := 0; x < 5; x++ {

		if contadoraux+1 <= x {
			_, _ = archivo.WriteString("struct" + strconv.Itoa(contadoraux) + "->" + "struct" + strconv.Itoa(contadoraux+1) + "\n")
			contadoraux++
		}
	}
	_, _ = archivo.WriteString("}" + "\n")
	var contadoaux2 int
	var contadoraux3 int
	var contadoaux5 int
	for i := iniciociclo; i < contadorGraficapart*5; i++ {

		contadoraux3 = contadoaux2
		var contador4 int
		a := linealizar[i].Frist

		_, _ = archivo.WriteString("subgraph cluster" + strconv.Itoa(i) + "{" + "\n")
		_, _ = archivo.WriteString("edge[dir=both]" + "\n")
		for a != nil {
			_, _ = archivo.WriteString("nodo" + strconv.Itoa(contadoaux2) + "[shape=record,label=\"{" + a.Store.Name + "|" + a.Store.Contact + "}\"];" + "\n")

			contadoaux2++
			contador4++
			a = a.Next
		}
		for x := 0; x < contador4; x++ {
			if contadoraux3+1+x < contadoaux2 {
				_, _ = archivo.WriteString("nodo" + strconv.Itoa(contadoraux3+x) + "->" + "nodo" + strconv.Itoa(contadoraux3+1+x) + "\n")

			}
		}

		_, _ = archivo.WriteString("}" + "\n")
		if contador4 > 0 {
			_, _ = archivo.WriteString("struct" + strconv.Itoa(contadoaux5) + "-> nodo" + strconv.Itoa(contadoraux3) + " [lhead=cluster" + strconv.Itoa(contadoaux5) + "];" + "\n")
		}
		contadoaux5++
	}

	_, _ = archivo.WriteString("}" + "\n")
	archivo.Close()
	iniciociclo = contadorGraficapart * 5

	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./graficoLinealizado"+strconv.Itoa(contadorGraficapart)+".dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("graficia"+strconv.Itoa(contadorGraficapart)+".pdf", cmd, os.FileMode(mode))
}

func guardar(w http.ResponseWriter, r *http.Request) {

	var jsongenerado vector
	var datosguardar datos

	jsongenerado.Datos = make([]datos, 0, len(indicesfor))
	var contador int
	var contador2 int
	for i := 0; i < len(indicesfor); i++ {

		var deparguardar []departamentos
		var deparindividual departamentos
		for j := 0; j < len(depa); j++ {
			contador2++
			var tiendasguardar []Store
			var tiendaindividual Store
			for k := contador; k < (contador2)*5; k++ {
				a := linealizar[k].Frist

				for a != nil {
					tiendaindividual = Store{a.Store.Name, a.Store.Description, a.Store.Contact, a.Store.Score}
					tiendasguardar = append(tiendasguardar, tiendaindividual)
					a = a.Next
				}
			}
			contador = (contador2) * 5
			deparindividual = departamentos{depa[j], tiendasguardar}
			deparguardar = append(deparguardar, deparindividual)
		}

		datosguardar = datos{indicesfor[i], deparguardar}
		jsongenerado.Datos = append(jsongenerado.Datos, datosguardar)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsongenerado)
	datos, _ := json.MarshalIndent(jsongenerado, "", " ")
	_ = ioutil.WriteFile("ArchivoGenerado.json", datos, 0644)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/guardar", guardar).Methods("GET")
	router.HandleFunc("/getArreglo", graficar).Methods("GET")
	router.HandleFunc("/cargartienda", agregar).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", busquedaEspecificaTienda).Methods("POST")
	router.HandleFunc("/id/{id}", busquedaPosicionLinealizada).Methods("GET")
	router.HandleFunc("/Eliminar", eliminarTienda).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))

}

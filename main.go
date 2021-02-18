package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"strconv"

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
	/*for i := 0; i < len(vectorJSON.Datos); i++ {
		fmt.Println("Indice: " + (vectorJSON.Datos[i].Indice))
		for j := 0; j < len(vectorJSON.Datos[i].Departamentos); j++ {
			fmt.Println("Departamento: " + vectorJSON.Datos[i].Departamentos[j].Departamento)
			for x := 0; x < len(vectorJSON.Datos[i].Departamentos[j].Tiendas); x++ {
				fmt.Println("Tienda: " + vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Name)
				fmt.Println("Contacto: " + vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Description)
				fmt.Println("Descripcion: " + vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Contact)
				fmt.Println("Calificacion: " + strconv.Itoa(vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Score))
			}
		}
	}*/
}

func linealizarMatriz() {
	linealizar = make([]estructura.List, 0, len(vectorJSON.Datos)*len(vectorJSON.Datos[0].Departamentos)*5)
	for i := 0; i < len(vectorJSON.Datos); i++ {
		fmt.Println("Indice: " + (vectorJSON.Datos[i].Indice))
		for j := 0; j < len(vectorJSON.Datos[i].Departamentos); j++ {
			fmt.Println("Departamento: " + vectorJSON.Datos[i].Departamentos[j].Departamento)
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
	}

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
	a := linealizar[posicion]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(a)
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
	var respuesta Store
	for i := 0; i < len(vectorJSON.Datos); i++ {
		for j := 0; j < len(vectorJSON.Datos[i].Departamentos); j++ {
			if varaux.Departamento == vectorJSON.Datos[i].Departamentos[j].Departamento {
				for x := 0; x < len(vectorJSON.Datos[i].Departamentos[j].Tiendas); x++ {
					if varaux.Nombre == vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Name && varaux.Calificacion == vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Score {
						respuesta.Name = vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Name
						respuesta.Description = vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Description
						respuesta.Contact = vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Contact
						respuesta.Score = vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Score
					}
				}
			}
		}
	}
	json.NewEncoder(w).Encode(respuesta)
}
func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Funcionan EDD")
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/agregar", agregar).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", busquedaEspecificaTienda).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./tienda"
	"github.com/gorilla/mux"
)

type vector struct {
	Datos []datos `json:"Datos"`
}

type datos struct {
	indice        string          `json:"Indice"`
	Departamentos []departamentos `json:"Departamentos"`
}

type departamentos struct {
	Departamento string         `json:"Nombre"`
	Tiendas      []tienda.Store `json:"Tiendas"`
}

func agregar(w http.ResponseWriter, r *http.Request) {
	var ms vector
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &ms)

	fmt.Println(ms.Datos)

	json.NewEncoder(w).Encode(ms)
}

func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Funcionan EDD")
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/agregar", agregar).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))

}

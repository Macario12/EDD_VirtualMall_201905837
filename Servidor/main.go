package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"strconv"

	"os"

	"./Grafo"
	"./Tabla_Hash"
	"./arbol_avl"
	"./arbol_b"
	"./estructura"
	"./matrizDispersa"
	"./pedidos"
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
	Logo        string `json:"Logo"`
}

type structaux struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}

var vectorJSON vector
var linealizar []estructura.List

func agregar(w http.ResponseWriter, r *http.Request) {
	fmt.Println(vectorJSON)
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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
				logoaux := vectorJSON.Datos[i].Departamentos[j].Tiendas[x].Logo

				arbolnuevo := arbol_avl.Newtree()
				arbolaux64 := "soy el imagen Base 64"
				pedidoaux := "Hola soy la img en 64"
				fechaaux := "-"
				fmt.Println(&arbolnuevo)
				tiendaaux := tienda.Store{nombreaux, descraux, cotactaux, scoreaux, logoaux, arbolnuevo, &arbolaux64, &pedidoaux, &fechaaux}
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
		tienda := tienda.Store{a.Store.Name, a.Store.Description, a.Store.Contact, a.Store.Score, a.Store.Logo, a.Store.Productos, a.Store.Arbol64, a.Store.Pedidos64, a.Store.Fecha}
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
		fmt.Println(a.Store.Logo)

		a = a.Next
	}
}

func codigoASCII(nombre string) int {

	var ascii int
	var inde int
	for i, c := range nombre {
		inde = i
		ascii += int(c)
	}
	fmt.Println(inde)
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

	for i := 0; i < len(indicesfor); i++ {
		for j := 0; j < len(depa); j++ {
			contadorGraficapart++
			graficodelArreglo(i, j, contadorGraficapart)
		}
	}

	var sliceTiendas []tienda.Store
	for i := 0; i < len(linealizar); i++ {
		listaObtenida := linealizar[i]
		a := listaObtenida.Frist
		for a != nil {
			tienda := tienda.Store{a.Store.Name, a.Store.Description, a.Store.Contact, a.Store.Score, a.Store.Logo, a.Store.Productos, a.Store.Arbol64, a.Store.Pedidos64, a.Store.Fecha}
			sliceTiendas = append(sliceTiendas, tienda)
			a = a.Next
		}
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(sliceTiendas)
	iniciociclo = 0
	posicionesvector = 0
}

var iniciociclo int
var posicionesvector int

func graficodelArreglo(i int, j int, contadorGraficapart int) {
	archivo, _ := os.Create("graficoLinealizado" + strconv.Itoa(contadorGraficapart) + ".dot")
	_, _ = archivo.WriteString("digraph grafico{" + "\n")
	_, _ = archivo.WriteString("compound=true;" + "\n")
	_, _ = archivo.WriteString("color=orangered" + "\n")
	_, _ = archivo.WriteString("subgraph cluster0{" + "\n")
	_, _ = archivo.WriteString("edge[minlen=0.1,color=\"greenyellow\", dir=fordware]" + "\n")
	var contador int
	var contadoraux int
	for x := 0; x < 5; x++ {

		_, _ = archivo.WriteString("struct" + strconv.Itoa(contador) + "[shape=record,color=\".7 .3 1.0\",label=\"" + (vectorJSON.Datos[i].Indice) + "|" + vectorJSON.Datos[i].Departamentos[j].Departamento + "|{" + strconv.Itoa(x+1) + "| pos:" + strconv.Itoa(posicionesvector) + "}\"];" + "\n")
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
		_, _ = archivo.WriteString("edge[color=\"steelblue\",dir=both]" + "\n")
		for a != nil {
			_, _ = archivo.WriteString("nodo" + strconv.Itoa(contadoaux2) + "[shape=record,color=\"darkturquoise\",label=\"{" + "{" + strconv.Itoa(codigoASCII(a.Store.Name)) + "|" + a.Store.Name + "}" + "|" + a.Store.Contact + "}\"];" + "\n")

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
					tiendaindividual = Store{a.Store.Name, a.Store.Description, a.Store.Contact, a.Store.Score, a.Store.Logo}
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

func AgregarInventario(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &arbol_avl.VectorInventario)
	agregarProductos()
	json.NewEncoder(w).Encode(arbol_avl.VectorInventario)
}

func agregarProductos() {
	for x := 0; x < len(arbol_avl.VectorInventario.Inventarios); x++ {
		indice := string(arbol_avl.VectorInventario.Inventarios[x].Tienda[0])

		var posicionL int
		var posi, posj int
		for i := 0; i < len(indicesfor); i++ {
			if indice == (indicesfor[i]) {
				posi = i
			}
		}
		for j := 0; j < len(depa); j++ {
			if arbol_avl.VectorInventario.Inventarios[x].Departamento == (depa[j]) {
				posj = j
			}
		}
		posicionL = ((posi*len(depa))+posj)*5 + (arbol_avl.VectorInventario.Inventarios[x].Calificacion - 1)
		fmt.Println(posicionL)
		//TRABAJANDO CON EL ARBOL QUE DEBE DE AGREGAR MUCHOS PORDUCTOS DESPUES DE UNA AGREGADA

		arbolaux := linealizar[posicionL].SearchStore(arbol_avl.VectorInventario.Inventarios[x].Tienda).GetProductos()

		fmt.Println(&arbolaux)
		for y := 0; y < len(arbol_avl.VectorInventario.Inventarios[x].Productoss); y++ {

			nombreaux := arbol_avl.VectorInventario.Inventarios[x].Productoss[y].Nombre
			codigoaux := arbol_avl.VectorInventario.Inventarios[x].Productoss[y].Codigo
			descripcionaux := arbol_avl.VectorInventario.Inventarios[x].Productoss[y].Descripcion
			precioaux := arbol_avl.VectorInventario.Inventarios[x].Productoss[y].Precio
			cantidadaux := arbol_avl.VectorInventario.Inventarios[x].Productoss[y].Cantidad
			imagenaux := arbol_avl.VectorInventario.Inventarios[x].Productoss[y].Imagen
			almacenamientoaux := arbol_avl.VectorInventario.Inventarios[x].Productoss[y].Almacenamiento

			productoaux := *arbol_avl.NewInventario(nombreaux, codigoaux, descripcionaux, precioaux, cantidadaux, imagenaux, almacenamientoaux)
			//fmt.Print(productoaux)
			arbolaux.Insertroot(productoaux)
		}

		arbol_avl.PreOrden(arbolaux.Root)

		fmt.Println("-----------------------------------------")

		*linealizar[posicionL].SearchStore(arbol_avl.VectorInventario.Inventarios[x].Tienda).Arbol64 = arbol_avl.Graficararbol(arbolaux.Root, arbol_avl.VectorInventario.Inventarios[x].Tienda+strconv.Itoa(x))
	}

}

var calendarioPedidos matrizDispersa.ListaDoble

type ContenedorAño struct {
	año   string
	meses matrizDispersa.ListaDoble
}

type ContenedorMes struct {
	mes    string
	matris *matrizDispersa.MatrisDispersa
}

func cargarPedidos(w http.ResponseWriter, r *http.Request) {
	var contenedorPedidos pedidos.PedidosP
	_ = json.NewDecoder(r.Body).Decode(&contenedorPedidos)

	//se agregan los años
	for _, pedido := range contenedorPedidos.PedidosPr {
		fechaString := strings.Split(pedido.Fecha, "-")
		año := fechaString[2]
		añoObject := ContenedorAño{año: año}
		añoAgregado := false
		for i := 0; i < calendarioPedidos.GetLen(); i++ {
			casillaAño := calendarioPedidos.Obtener(i).(ContenedorAño)
			if año == casillaAño.año {
				añoAgregado = true
				break
			}
		}

		//se agregan los meses
		for _, pedido_2 := range contenedorPedidos.PedidosPr {
			fechaString_2 := strings.Split(pedido_2.Fecha, "-")
			año_2 := fechaString_2[2]
			mes := fechaString_2[1]
			if año_2 == año {
				mesObject := ContenedorMes{mes: mes}
				mesAgregado := false
				for i := 0; i < añoObject.meses.GetLen(); i++ {
					casillaMes := añoObject.meses.Obtener(i).(ContenedorMes)
					if mes == casillaMes.mes {
						mesAgregado = true
						break
					}
				}
				if mesAgregado == false {
					matris := new(matrizDispersa.MatrisDispersa)
					mesObject.matris = matris
					añoObject.meses.Agregar(mesObject)
				}
			}
		}
		if añoAgregado == false {
			calendarioPedidos.Agregar(añoObject)
		}
	}

	//agregar los pedidos
	for _, pedido := range contenedorPedidos.PedidosPr {
		fechaString := strings.Split(pedido.Fecha, "-")
		año := fechaString[2]
		mes := fechaString[1]
		dia := fechaString[0]
		for a := 0; a < calendarioPedidos.GetLen(); a++ {
			contAño := calendarioPedidos.Obtener(a).(ContenedorAño)
			for b := 0; b < contAño.meses.GetLen(); b++ {
				contMes := contAño.meses.Obtener(b).(ContenedorMes)
				if contAño.año == año && contMes.mes == mes {
					contMes.matris.Insertar(pedido, pedido.Departamento, dia)
				}
			}
		}
	}

	for x := 0; x < len(contenedorPedidos.PedidosPr); x++ {
		indice := string(contenedorPedidos.PedidosPr[x].Tienda[0])

		var posicionL int
		var posi, posj int
		for i := 0; i < len(indicesfor); i++ {
			if indice == (indicesfor[i]) {
				posi = i
			}
		}
		for j := 0; j < len(depa); j++ {
			if contenedorPedidos.PedidosPr[x].Departamento == (depa[j]) {
				posj = j
			}
		}
		posicionL = ((posi*len(depa))+posj)*5 + (contenedorPedidos.PedidosPr[x].Calificacion - 1)
		*linealizar[posicionL].SearchStore(contenedorPedidos.PedidosPr[x].Tienda).Fecha = contenedorPedidos.PedidosPr[x].Fecha
		fmt.Println(posicionL)
		for i := 0; i < calendarioPedidos.GetLen(); i++ {
			contAño := calendarioPedidos.Obtener(i).(ContenedorAño)
			for a := 0; a < contAño.meses.GetLen(); a++ {
				contMes := contAño.meses.Obtener(a).(ContenedorMes)
				*linealizar[posicionL].SearchStore(contenedorPedidos.PedidosPr[x].Tienda).Pedidos64 = contMes.matris.Graficar(contAño.año + "_" + contMes.mes)
			}
		}

	}

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contenedorPedidos)
}

//Usuario
func cargarUsuarios(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println("CARGA DE USUARIOS")
	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &arbol_b.VectorUsers)
	agregarUsuario()
	json.NewEncoder(w).Encode(arbol_b.VectorUsers)
}

// LLAVE PENDIENTE
func cargarllave(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	fmt.Println("LLLAVE")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &arbol_b.Llave)
	fmt.Println(arbol_b.Llave)
	json.NewEncoder(w).Encode(arbol_b.Llave)
}

var arbol *arbol_b.Arbol

//Grafo
func graficarGrafo(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	recorrido := grafo.ObtenerRecorrido(grafo.EstadoInicial, grafo.EstadoFinal)
	json.NewEncoder(w).Encode(grafo.Graficar("GRafoCa", recorrido.Trayectoria, recorrido.EstadosTrayectoria))
}

func agregarUsuario() {

	for x := 0; x < len(arbol_b.VectorUsers.Usuarios); x++ {
		encriptar := []byte(arbol_b.VectorUsers.Usuarios[x].Password)
		hash := sha256.Sum256(encriptar)
		usuario := arbol_b.NewUser(arbol_b.VectorUsers.Usuarios[x].DPI, arbol_b.VectorUsers.Usuarios[x].Nombre, arbol_b.VectorUsers.Usuarios[x].Correo, hex.EncodeToString(hash[:]), arbol_b.VectorUsers.Usuarios[x].Cuenta)

		arbol.Insertar(arbol_b.NewKey(*usuario))

	}

}

//Cargar Unico Usuario
func cargarUsuario(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println("CARGA DE Usuario")
	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &usuariounico)
	agregarUsuarioUnico()
	json.NewEncoder(w).Encode(usuariounico)
}

var usuariounico arbol_b.User

func agregarUsuarioUnico() {

	usuario := arbol_b.NewUser(usuariounico.DPI, usuariounico.Nombre, usuariounico.Correo, usuariounico.Password, usuariounico.Cuenta)

	arbol.Insertar(arbol_b.NewKey(*usuario))

}

var usuariob arbol_b.User

//Busqueda de usuario pa Login

func buscarUsuario(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println("CARGA DE Usuario")
	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &usuariob)

	json.NewEncoder(w).Encode(arbol_b.Buscar(arbol.Raiz, usuariob, nil, 0))
}

var sliceProducots []arbol_avl.Inventario

func PreOrden(n *arbol_avl.Node) {
	if n != nil {
		productoR := arbol_avl.Inventario{n.Data.Nombre, n.Data.Codigo, n.Data.Descripcion, n.Data.Precio, n.Data.Cantidad, n.Data.Imagen, n.Data.Almacenamiento}
		sliceProducots = append(sliceProducots, productoR)
		fmt.Println(n.Data.Nombre)
		PreOrden(n.Left)
		PreOrden(n.Right)
	}
}

type StoreReturn struct {
	Departamento string
	Name         string `json:"Nombre"`
	Description  string `json:"Descripcion"`
	Contact      string `json:"Contacto"`
	Score        int    `json:"Calificacion"`
	Logo         string `json:"Logo"`
	Productos    []arbol_avl.Inventario
	Raiz         *arbol_avl.Tree
	Arbol64      *string
	Pedidos64    *string
	Fecha        *string
}

//Retorno tiendas
func retornarTiendas(w http.ResponseWriter, r *http.Request) {

	var sliceTiendas []StoreReturn
	var contadorrer int
	var departamentoaux string
	var contadorlinealizarau int
	for i := 0; i < len(linealizar); i++ {
		listaObtenida := linealizar[i]
		a := listaObtenida.Frist
		departamentoaux = depa[contadorrer]
		contadorlinealizarau++
		if contadorlinealizarau == 5 {
			contadorlinealizarau = 0
			contadorrer++
		}
		if contadorrer == len(depa) {
			contadorrer = 0
		}
		for a != nil {

			producto := a.Store.Productos.Root

			PreOrden(producto)
			fmt.Println(sliceProducots)

			tienda := StoreReturn{departamentoaux, a.Store.Name, a.Store.Description, a.Store.Contact, a.Store.Score, a.Store.Logo, sliceProducots, a.Store.Productos, a.Store.Arbol64, a.Store.Pedidos64, a.Store.Fecha}

			sliceTiendas = append(sliceTiendas, tienda)

			sliceProducots = nil
			a = a.Next
		}
	}

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(sliceTiendas)
}

//Grafos de Arboles, Sin, Cifrado, Sencible

func graficarArbolB(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(arbol.GraficarArbolSin())
}

func graficarArbolBCifrado(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(arbol.GraficarArbolCifrado())
}

func graficarArbolBCifradoSencible(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(arbol.GraficarArbolCifradoSencible())

}

//GRAFO

//Cargar Grafo

func cargarGrafo(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println("CARGA DE GRAFO")
	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &Grafo.Grafojson)
	grafocrear()
	json.NewEncoder(w).Encode(Grafo.Grafojson)
}

var grafo Grafo.Grafo

func grafocrear() {
	var enlaces []Grafo.Enlace
	for _, nodo := range Grafo.Grafojson.Nodos {
		for _, enlace := range nodo.EnlancesA {
			enlaces = append(enlaces, Grafo.Enlace{
				EstadoInicial:   nodo.Nombre,
				EstadoFinal:     enlace.Nombre,
				PesoDeRecorrido: float64(enlace.Distancia),
			})

			enlaces = append(enlaces, Grafo.Enlace{
				EstadoInicial:   enlace.Nombre,
				EstadoFinal:     nodo.Nombre,
				PesoDeRecorrido: float64(enlace.Distancia),
			})
		}

	}

	grafo = Grafo.Grafo{
		Enlaces:       enlaces,
		EstadoInicial: Grafo.Grafojson.PosicionInicialRobot,
		EstadoFinal:   Grafo.Grafojson.Entrega,
	}

	recorrido := grafo.ObtenerRecorrido(grafo.EstadoInicial, grafo.EstadoFinal)

	grafo.Graficar("Grafo", recorrido.Trayectoria, recorrido.EstadosTrayectoria)
}

type Peticion struct {
	Tienda     tienda.Store `json:"Tienda"`
	Usuario    arbol_b.User `json:"Usuario"`
	IdTabla    string       `json:"IdTabla"`
	Comentario string       `json:"comentario"`
	Fecha      string       `json:"Fecha"`
}

func agregarComentarioT(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	fmt.Println("Subir Comentario")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	var peticion Peticion
	json.Unmarshal(reqBody, &peticion)
	fmt.Println("id tabla: " + peticion.IdTabla)
	fmt.Println("Comentario: " + peticion.Comentario)
	for i := 0; i < len(linealizar); i++ {
		nodo := linealizar[i].SearchNodo(peticion.Tienda.Name, peticion.Tienda.Contact)
		if nodo != nil {	
			tabla := nodo.GetTabla()
			if peticion.IdTabla != "" {
				tabla = nodo.GetTabla().GetTabla(peticion.IdTabla)
				tabla.Insertar(peticion.Usuario.DPI, tabla.NewRegistro(&peticion.Usuario, peticion.Comentario))
			} else {
				tabla = nodo.GetTabla()
				registro := tabla.NewRegistro(&peticion.Usuario, peticion.Comentario)
				tabla.Insertar(peticion.Usuario.DPI, registro)
			}
			fmt.Println(tabla.GetAsList())
			json.NewEncoder(w).Encode(tabla.GetAsList())
		}
	}

	
	
}

func retornarComenTienda(w http.ResponseWriter, r *http.Request) {
	type Peticion struct {
		Tienda tienda.Store `json:"Tienda"`
	}
	type Respuesta struct {
		idTabla     string
		comentarios []*Tabla_Hash.Registro
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	fmt.Println("Subir Comentario")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	var peticion Peticion
	json.Unmarshal(reqBody, &peticion)
	_ = json.NewDecoder(r.Body).Decode(&peticion)
	for i := 0; i < len(linealizar); i++ {
		nodo := linealizar[i].SearchNodo(peticion.Tienda.Name, peticion.Tienda.Contact)
			if nodo != nil {
				fmt.Println("")
				fmt.Println(nodo.Comentarios.GetAsList())
				json.NewEncoder(w).Encode(nodo.Comentarios.GetAsList())
			}
		
	}

	fmt.Println("Retornando COment")
}

func subirComentarioProducto(w http.ResponseWriter, r *http.Request) {
	type Peticion struct {
		Producto   arbol_avl.Inventario `json:"Producto"`
		Usuario    arbol_b.User `json:"Usuario"`
		IdTabla    string             `json:"IdTabla"`
		Comentario string             `json:"Comentario"`
	}
/*	var peticion Peticion
	_ = json.NewDecoder(r.Body).Decode(&peticion)
	fmt.Println("id tabla: " + peticion.IdTabla)
	for _, listaDoble := range matrisLinearizada {
		for a := 0; a < listaDoble.GetLen(); a++ {
			nodo := listaDoble.ObtenerNodo(a)
			producto := nodo.Inventario.GetNodo(peticion.Producto.Codigo)
			if producto != nil {
				var tabla *estructuras.TablaHash
				if peticion.IdTabla != "" {
					tabla = producto.Comentarios.GetTabla(peticion.IdTabla)
					tabla.Insertar(peticion.Usuario.DPI, tabla.NewRegistro(&peticion.Usuario, peticion.Comentario))
				} else {
					tabla = producto.Comentarios
					tabla.Insertar(peticion.Usuario.DPI, tabla.NewRegistro(&peticion.Usuario, peticion.Comentario))
				}
				json.NewEncoder(w).Encode(tabla.GetAsList())
			}
		}
	}

	*/
}

//MAIN
func main() {
	arbol = arbol_b.NewArbol(5)
	encriptar := []byte("1234")
	hash := sha256.Sum256(encriptar)
	userDefault := arbol_b.NewUser(1234567890101, "EDD2021", " auxiliar@edd.com", hex.EncodeToString(hash[:]), "Admin")
	arbol.Insertar(arbol_b.NewKey(*userDefault))
	router := mux.NewRouter()

	router.HandleFunc("/tiendascargadas", retornarTiendas).Methods("GET")
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/guardar", guardar).Methods("GET")
	router.HandleFunc("/getArreglo", graficar).Methods("GET")
	router.HandleFunc("/cargartienda", agregar).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", busquedaEspecificaTienda).Methods("POST")
	router.HandleFunc("/id/{id}", busquedaPosicionLinealizada).Methods("GET")
	router.HandleFunc("/Eliminar", eliminarTienda).Methods("DELETE")
	router.HandleFunc("/agregarINV", AgregarInventario).Methods("POST")
	router.HandleFunc("/usuarios", cargarUsuarios).Methods("POST")
	router.HandleFunc("/crearUsuario", cargarUsuario).Methods("POST")
	router.HandleFunc("/cargarPedidos", cargarPedidos).Methods("POST")
	router.HandleFunc("/buscarUsuario", buscarUsuario).Methods("POST")
	router.HandleFunc("/llave", cargarllave).Methods("POST")
	router.HandleFunc("/arbolS", graficarArbolB).Methods("GET")
	router.HandleFunc("/arbolC", graficarArbolBCifrado).Methods("GET")

	router.HandleFunc("/grafo", graficarGrafo).Methods("GET")
	router.HandleFunc("/arbolCS", graficarArbolBCifradoSencible).Methods("GET")
	router.HandleFunc("/Cgrafo", cargarGrafo).Methods("POST")

	//COMENTARIOS

	router.HandleFunc("/getComentariosTienda", retornarComenTienda).Methods("POST")
	router.HandleFunc("/subirComentarioTienda", agregarComentarioT).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))

}

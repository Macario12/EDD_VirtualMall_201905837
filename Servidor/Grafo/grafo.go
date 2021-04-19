package Grafo

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)


type Grafo struct {
	Enlaces       []Enlace 
	EstadoInicial string
	EstadoFinal   string
}

type Enlace struct {
	EstadoInicial   string
	EstadoFinal     string
	PesoDeRecorrido float64
}

type Union struct {
	EstadoAnterior      string
	EstadoActual        string
	RecorridoAcumulado  []string
	stringRecorridoAcum string
	PesoAcumulado       float64
	SeguirEvaluando     bool
}

type Recorrido struct {
	EstadosTrayectoria []string
	Trayectoria        string
	PesoAcumulado      float64
}

func (g *Grafo) ObtenerRecorrido(estadoInicial string, estadoFinal string) Recorrido {
	if estadoInicial == estadoFinal {
		return Recorrido{}
	}
	var arbol [][]Union
	var nivelCero []Union
	nivelCero = append(nivelCero, Union{
		EstadoAnterior:      estadoInicial,
		EstadoActual:        estadoInicial,
		RecorridoAcumulado:  []string{estadoInicial},
		stringRecorridoAcum: estadoInicial,
		SeguirEvaluando:     true,
	})
	arbol = append(arbol, nivelCero)

	for continuar(arbol[len(arbol)-1], estadoFinal) {
		var nivel []Union
		for _, union := range arbol[len(arbol)-1] {
			nivel = append(nivel, union.getUnionsOfUnion(g.Enlaces, estadoFinal)...)
		}
		arbol = append(arbol, nivel)
	}

	recorrido := ordenar(arbol[len(arbol)-1])[0]
	return Recorrido{recorrido.RecorridoAcumulado, recorrido.stringRecorridoAcum, recorrido.PesoAcumulado}
}

func continuar(ultimoNivel []Union, estadoFinal string) bool {
	for _, union := range ultimoNivel {
		if union.EstadoActual != estadoFinal {
			return true
		}
	}
	return false
}

func (u Union) getUnionsOfUnion(enlaces []Enlace, estadoFinal string) []Union {
	var uniones []Union
	if u.SeguirEvaluando {
		for _, enlace := range enlaces {
			if enlace.EstadoInicial == u.EstadoActual &&
				!recorridoAcumContineEstado(u.RecorridoAcumulado, enlace.EstadoFinal) {
				union := Union{}
				union.EstadoAnterior = u.EstadoActual
				union.EstadoActual = enlace.EstadoFinal
				var list []string
				list = append(list, u.RecorridoAcumulado...)
				union.RecorridoAcumulado = append(list, enlace.EstadoFinal)
				s := fmt.Sprintf("%v", enlace.PesoDeRecorrido)
				union.stringRecorridoAcum += u.stringRecorridoAcum + " -" + s + "→ " + union.EstadoActual
				union.PesoAcumulado = u.PesoAcumulado + enlace.PesoDeRecorrido
				if union.EstadoActual != estadoFinal {
					union.SeguirEvaluando = true
				}
				uniones = append(uniones, union)
			}
		}
	} else {
		uniones = append(uniones, u)
	}
	return uniones
}

func recorridoAcumContineEstado(recorrido []string, estado string) bool {
	for _, e := range recorrido {
		if e == estado {
			return true
		}
	}
	return false
}

func ordenar(nivel []Union) []Union {
	for i := 0; i < len(nivel); i++ {
		for j := 0; j < len(nivel)-i-1; j++ {
			if nivel[j].PesoAcumulado > nivel[j+1].PesoAcumulado {
				nivel[j], nivel[j+1] = nivel[j+1], nivel[j]
			}
		}
	}
	return nivel
}

func contiene(estados []string, estado string) bool {
	for _, e := range estados {
		if e == estado {
			return true
		}
	}
	return false
}

func contieneEnlace(enlaces []Enlace, enlace Enlace) bool {
	for _, e := range enlaces {
		if e.EstadoInicial == enlace.EstadoInicial && e.EstadoFinal == enlace.EstadoFinal ||
			e.EstadoFinal == enlace.EstadoInicial && e.EstadoInicial == enlace.EstadoFinal {
			return true
		}
	}
	return false
}

func format(texto string) string {
	return strings.ReplaceAll(texto, " ", "_")
}

func (g *Grafo) Graficar(nombreGrafica, info string, estadosPintar []string) string {
	texto := "digraph G { \n"
	texto += "layout=\"circo\" \n"
	texto += "node [] \n"

	texto += "subgraph cluster_0 { color = gray100;"
	texto += "info[label=\"" + info + "\" shape=box color=\"steelblue\"] \n"
	texto += "}"

	texto += "subgraph cluster_1 { color = gray100;"
	texto += format(g.EstadoInicial) + "[label=\"" + g.EstadoInicial + "\" color=\"coral\"] \n"
	texto += format(g.EstadoFinal) + "[label=\"" + g.EstadoFinal + "\"  color=\"coral\"] \n"

	//Pinto los nodos del recorrido
	var estados []string
	estados = append(estados, g.EstadoInicial)
	estados = append(estados, g.EstadoFinal)
	for _, estado := range estadosPintar {
		if !contiene(estados, estado) {
			texto += format(estado) + "[label=\"" + estado + "\" color=greenyellow] \n"
		}
		estados = append(estados, estado)
	}
	//Pinto los nodos que no son del recorrido
	for _, enlace := range g.Enlaces {
		if !contiene(estados, enlace.EstadoInicial) {
			texto += format(enlace.EstadoInicial) + "[label=\"" + enlace.EstadoInicial + "\" color=firebrick4] \n"
			estados = append(estados, enlace.EstadoInicial)
		}
		if !contiene(estados, enlace.EstadoFinal) {
			texto += format(enlace.EstadoFinal) + "[label=\"" + enlace.EstadoFinal +"\" color=firebrick4] \n"
			estados = append(estados, enlace.EstadoFinal)
		}
	}
	//Pinto los enlaces del recorrido
	var enlaces []Enlace
	for i := 0; i < len(estadosPintar)-1; i++ {
		if estadosPintar[i] != estadosPintar[i+1] {
			for _, e := range g.Enlaces {
				if (e.EstadoInicial == estadosPintar[i] && e.EstadoFinal == estadosPintar[i+1] ||
					e.EstadoInicial == estadosPintar[i+1] && e.EstadoFinal == estadosPintar[i]) &&
					!contieneEnlace(enlaces, e) {
					s := fmt.Sprintf("%v", e.PesoDeRecorrido)
					texto += format(e.EstadoInicial) + " -> " + format(e.EstadoFinal) + " [label=\"" + s + "\" dir=both color=greenyellow] \n"
					enlaces = append(enlaces, e)
				}
			}
		}
	}
	//Pinto los enlaces que no son del recorrido
	for _, e := range g.Enlaces {
		if !contieneEnlace(enlaces, e) {
			s := fmt.Sprintf("%v", e.PesoDeRecorrido)
			texto += format(e.EstadoInicial) + " -> " + format(e.EstadoFinal) + " [label=\"" + s + "\" dir=both color=\".7 .3 1.0\"] \n"
			enlaces = append(enlaces, e)
		}
	}

	texto += "}"
	texto += "info -> " + format(g.EstadoInicial) + " [color=\"firebrick4\"] \n"
	texto += "}"

	_ = ioutil.WriteFile(nombreGrafica+".dot", []byte(texto), 0644)
	cmd := exec.Command("dot", "-Tpng", nombreGrafica+".dot", "-o", "./"+nombreGrafica+".png")
	_ = cmd.Run()

	file, _ := os.Open("./" + nombreGrafica + ".png")
	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded
}

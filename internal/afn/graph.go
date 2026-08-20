package afn

import (
	"fmt"
	"strings"
)

// Genera representación DOT para el AFN
func GenerarDOT(afn *AFN) string {
	var builder strings.Builder
	builder.WriteString("digraph AFN {\n")
	builder.WriteString("\trankdir=LR;\n")
	builder.WriteString("\tnode [shape=circle, fontname=\"Arial\"];\n")
	
	// Estado de aceptación
	builder.WriteString(fmt.Sprintf("\tnode%d [shape=doublecircle];\n", afn.Final.ID))

	// Indicador de estado inicial
	builder.WriteString("\tinicio [shape=point];\n")
	builder.WriteString(fmt.Sprintf("\tinicio -> node%d;\n", afn.Inicial.ID))

	visitados := make(map[int]bool)
	cola := []*Estado{afn.Inicial}
	visitados[afn.Inicial.ID] = true

	// BFS para recorrer el AFN
	for len(cola) > 0 {
		actual := cola[0]
		cola = cola[1:]

		for _, t := range actual.Transiciones {
			label := t.Simbolo
			if label == "|" {
				label = "\\|" // Escape para Graphviz
			}
			
			builder.WriteString(fmt.Sprintf("\tnode%d -> node%d [label=\"%s\"];\n", actual.ID, t.Destino.ID, label))

			if !visitados[t.Destino.ID] {
				visitados[t.Destino.ID] = true
				cola = append(cola, t.Destino)
			}
		}
	}

	builder.WriteString("}\n")
	return builder.String()
}
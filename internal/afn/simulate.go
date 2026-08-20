package afn

// Calcula el Cierre-Épsilon de un conjunto de estados
func CerraduraEpsilon(estados []*Estado) []*Estado {
	pila := make([]*Estado, len(estados))
	copy(pila, estados)
	visitados := make(map[int]bool)
	var resultado []*Estado

	for _, e := range estados {
		visitados[e.ID] = true
		resultado = append(resultado, e)
	}

	for len(pila) > 0 {
		actual := pila[len(pila)-1]
		pila = pila[:len(pila)-1]

		for _, t := range actual.Transiciones {
			if t.Simbolo == "ε" {
				if !visitados[t.Destino.ID] {
					visitados[t.Destino.ID] = true
					resultado = append(resultado, t.Destino)
					pila = append(pila, t.Destino)
				}
			}
		}
	}
	return resultado
}

// Retorna los estados alcanzables al consumir un símbolo específico
func Mover(estados []*Estado, simbolo string) []*Estado {
	var resultado []*Estado
	visitados := make(map[int]bool)

	for _, e := range estados {
		for _, t := range e.Transiciones {
			if t.Simbolo == simbolo {
				if !visitados[t.Destino.ID] {
					visitados[t.Destino.ID] = true
					resultado = append(resultado, t.Destino)
				}
			}
		}
	}
	return resultado
}

// Simula la cadena en el AFN
func Simular(afn *AFN, cadena string) bool {
	estadosActuales := CerraduraEpsilon([]*Estado{afn.Inicial})

	for _, char := range cadena {
		simbolo := string(char)
		estadosActuales = CerraduraEpsilon(Mover(estadosActuales, simbolo))
	}

	for _, e := range estadosActuales {
		if e.ID == afn.Final.ID {
			return true
		}
	}
	return false
}
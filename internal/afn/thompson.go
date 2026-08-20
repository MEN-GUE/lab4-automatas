package afn

import "lab4/internal/ast"

type Transicion struct {
	Simbolo string
	Destino *Estado
}

type Estado struct {
	ID           int
	Transiciones []Transicion
}

type AFN struct {
	Inicial *Estado
	Final   *Estado
}

var contadorEstados int

func ReiniciarContador() {
	contadorEstados = 0
}

func nuevoEstado() *Estado {
	estado := &Estado{ID: contadorEstados}
	contadorEstados++
	return estado
}

func ConstruirThompson(nodo *ast.Nodo) *AFN {
	if nodo == nil {
		return nil
	}

	switch nodo.Valor {
	case "~": // Concatenación
		afnIzq := ConstruirThompson(nodo.Izq)
		afnDer := ConstruirThompson(nodo.Der)
		// Transición épsilon del final del izquierdo al inicio del derecho
		afnIzq.Final.Transiciones = append(afnIzq.Final.Transiciones, Transicion{Simbolo: "ε", Destino: afnDer.Inicial})
		return &AFN{Inicial: afnIzq.Inicial, Final: afnDer.Final}

	case "|": // Unión
		afnIzq := ConstruirThompson(nodo.Izq)
		afnDer := ConstruirThompson(nodo.Der)
		inicial := nuevoEstado()
		final := nuevoEstado()

		inicial.Transiciones = append(inicial.Transiciones, Transicion{Simbolo: "ε", Destino: afnIzq.Inicial})
		inicial.Transiciones = append(inicial.Transiciones, Transicion{Simbolo: "ε", Destino: afnDer.Inicial})
		
		afnIzq.Final.Transiciones = append(afnIzq.Final.Transiciones, Transicion{Simbolo: "ε", Destino: final})
		afnDer.Final.Transiciones = append(afnDer.Final.Transiciones, Transicion{Simbolo: "ε", Destino: final})
		
		return &AFN{Inicial: inicial, Final: final}

	case "*": // Cerradura de Kleene
		afnBase := ConstruirThompson(nodo.Izq)
		inicial := nuevoEstado()
		final := nuevoEstado()

		inicial.Transiciones = append(inicial.Transiciones, Transicion{Simbolo: "ε", Destino: afnBase.Inicial})
		inicial.Transiciones = append(inicial.Transiciones, Transicion{Simbolo: "ε", Destino: final})
		
		afnBase.Final.Transiciones = append(afnBase.Final.Transiciones, Transicion{Simbolo: "ε", Destino: afnBase.Inicial})
		afnBase.Final.Transiciones = append(afnBase.Final.Transiciones, Transicion{Simbolo: "ε", Destino: final})
		
		return &AFN{Inicial: inicial, Final: final}

	default: // Caso base (Símbolo o ε)
		inicial := nuevoEstado()
		final := nuevoEstado()
		inicial.Transiciones = append(inicial.Transiciones, Transicion{Simbolo: nodo.Valor, Destino: final})
		return &AFN{Inicial: inicial, Final: final}
	}
}
package ast

type Nodo struct {
	Valor string
	Izq   *Nodo
	Der   *Nodo
}

type PilaNodos []*Nodo

func (p *PilaNodos) Push(v *Nodo) { *p = append(*p, v) }
func (p *PilaNodos) Pop() (*Nodo, bool) {
	if len(*p) == 0 { return nil, false }
	idx := len(*p) - 1
	val := (*p)[idx]
	*p = (*p)[:idx]
	return val, true
}

func ConstruirAST(postfix []string) *Nodo {
	var pila PilaNodos

	for _, token := range postfix {
		if token == "*" {
			n1, _ := pila.Pop()
			nodo := &Nodo{Valor: token, Izq: n1}
			pila.Push(nodo)
		} else if token == "|" || token == "~" {
			n2, _ := pila.Pop()
			n1, _ := pila.Pop()
			nodo := &Nodo{Valor: token, Izq: n1, Der: n2}
			pila.Push(nodo)
		} else {
			nodo := &Nodo{Valor: token}
			pila.Push(nodo)
		}
	}
	raiz, _ := pila.Pop()
	return raiz
}
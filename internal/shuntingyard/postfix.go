package shuntingyard

type PilaStrings []string

func (p *PilaStrings) Push(v string) { *p = append(*p, v) }
func (p *PilaStrings) Pop() (string, bool) {
	if len(*p) == 0 { return "", false }
	index := len(*p) - 1
	val := (*p)[index]
	*p = (*p)[:index]
	return val, true
}
func (p *PilaStrings) Peek() string {
	if len(*p) == 0 { return "" }
	return (*p)[len(*p)-1]
}
func (p *PilaStrings) IsEmpty() bool { return len(*p) == 0 }

func Tokenizar(regex string) []string {
	var tokens []string
	for _, c := range regex {
		tokens = append(tokens, string(c))
	}
	return tokens
}

func obtenerOperandoAnterior(res []string, fin int) ([]string, int) {
	if len(res) == 0 {
		return []string{}, 0
	}
	if res[fin] == ")" {
		pares := 1
		inicio := fin - 1
		for inicio >= 0 && pares > 0 {
			if res[inicio] == ")" { pares++ }
			if res[inicio] == "(" { pares-- }
			inicio--
		}
		return append([]string(nil), res[inicio+1:fin+1]...), inicio + 1
	}
	return []string{res[fin]}, fin
}

func SimplificarExtensiones(tokens []string) []string {
	var res []string
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if t == "+" {
			op, _ := obtenerOperandoAnterior(res, len(res)-1)
			res = append(res, op...)
			res = append(res, "*")
		} else if t == "?" {
			op, inicio := obtenerOperandoAnterior(res, len(res)-1)
			res = res[:inicio]
			res = append(res, "(")
			res = append(res, op...)
			res = append(res, "|", "ε", ")")
		} else {
			res = append(res, t)
		}
	}
	return res
}

func AgregarConcatenacionExplicita(tokens []string) []string {
	var res []string
	for i := 0; i < len(tokens); i++ {
		t1 := tokens[i]
		res = append(res, t1)
		if i+1 < len(tokens) {
			t2 := tokens[i+1]
			esT1Abierto := t1 == "(" || t1 == "|"
			esT2Cerrado := t2 == ")" || t2 == "|" || t2 == "*" || t2 == "+" || t2 == "?"
			if !esT1Abierto && !esT2Cerrado {
				res = append(res, "~")
			}
		}
	}
	return res
}

func Precedencia(t string) int {
	switch t {
	case "(": return 1
	case "|": return 2
	case "~": return 3
	case "?", "*", "+": return 4
	default: return 0
	}
}

func Convertir(tokens []string) []string {
	var postfix []string
	var pila PilaStrings

	for _, t := range tokens {
		switch t {
		case "(":
			pila.Push(t)
		case ")":
			for !pila.IsEmpty() && pila.Peek() != "(" {
				op, _ := pila.Pop()
				postfix = append(postfix, op)
			}
			pila.Pop()
		case "|", "~", "*":
			for !pila.IsEmpty() {
				if Precedencia(pila.Peek()) >= Precedencia(t) {
					op, _ := pila.Pop()
					postfix = append(postfix, op)
				} else {
					break
				}
			}
			pila.Push(t)
		default:
			postfix = append(postfix, t)
		}
	}
	for !pila.IsEmpty() {
		op, _ := pila.Pop()
		postfix = append(postfix, op)
	}
	return postfix
}
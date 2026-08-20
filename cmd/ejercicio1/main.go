package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"lab4/internal/afn"
	"lab4/internal/ast"
	"lab4/internal/shuntingyard"
)

func procesarLinea(linea string, index int) {
	partes := strings.Split(linea, " ")
	if len(partes) != 2 {
		fmt.Printf("Error: La línea %d no cumple con el formato 'regex cadena'\n", index)
		return
	}

	regex := partes[0]
	cadena := partes[1]

	fmt.Printf("\n================================================================================\n")
	fmt.Printf("Expresión Regular: %s\n", regex)
	fmt.Printf("Cadena a evaluar (w): %s\n", cadena)

	tokens := shuntingyard.Tokenizar(regex)
	tokensSimples := shuntingyard.SimplificarExtensiones(tokens)
	tokensConConcat := shuntingyard.AgregarConcatenacionExplicita(tokensSimples)
	postfix := shuntingyard.Convertir(tokensConConcat)

	// Creación de AST
	raizAST := ast.ConstruirAST(postfix)

	// Creación de AFN (Algoritmo de Thompson)
	afn.ReiniciarContador()
	automata := afn.ConstruirThompson(raizAST)

	// Graficación del AFN
	dotSource := afn.GenerarDOT(automata)
	dotFilename := fmt.Sprintf("afn_%d.dot", index)
	pngFilename := fmt.Sprintf("afn_%d.png", index)

	os.WriteFile(dotFilename, []byte(dotSource), 0644)
	cmd := exec.Command("dot", "-Tpng", dotFilename, "-o", pngFilename)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error al generar imagen con Graphviz:", err)
		return
	}

	fmt.Printf("✅ AFN generado exitosamente en: %s\n", pngFilename)
	exec.Command("xdg-open", pngFilename).Start()

	// Simulación
	pertenece := afn.Simular(automata, cadena)
	
	// Formato de salida de pertenencia (requerido)
	if pertenece {
		fmt.Printf("Resultado simulación: sí, '%s' pertenece a L(r)\n", cadena)
	} else {
		fmt.Printf("Resultado simulación: no, '%s' NO pertenece a L(r)\n", cadena)
	}
}

func main() {
	fmt.Println("=== SIMULADOR DE AUTÓMATAS (THOMPSON) ===")

	archivo, err := os.Open("expresiones.txt")
	if err != nil {
		fmt.Println("No se pudo abrir 'expresiones.txt'.")
		return
	}
	defer archivo.Close()

	escaner := bufio.NewScanner(archivo)
	index := 1

	for escaner.Scan() {
		linea := strings.TrimSpace(escaner.Text())
		if len(linea) == 0 {
			continue
		}
		procesarLinea(linea, index)
		index++
	}

	if err := escaner.Err(); err != nil {
		fmt.Println("Error leyendo el archivo:", err)
	}
}
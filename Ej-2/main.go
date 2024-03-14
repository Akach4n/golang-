package main

import "fmt"

func main() {
	// Definir dos variables para almacenar dos números
	var num1, num2 int
	// Pedir al usuario que ingrese dos números
	fmt.Print("Dame un número: ")
	fmt.Scan(&num1)
	fmt.Print("Dame otro número: ")
	fmt.Scan(&num2)
	// Llamar a una función para encontrar el máximo común divisor
	resultado := encontrarMCD(num1, num2)
	// Imprimir el máximo común divisor
	fmt.Println(resultado)
}

// Función para encontrar el máximo común divisor utilizando el algoritmo de Euclides
func encontrarMCD(a, b int) int {
	// Implementar el algoritmo de Euclides para encontrar el MCD
	var aux int
	for b != 0 {
		aux = b
		b = a % b
		a = aux

	}
	return a
}

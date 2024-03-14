package main

import "fmt"

func main() {
	// Define una lista de números
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Inicializa una variable para almacenar la suma de los números pares
	var total int
	// Itera sobre los números y suma los números pares
	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			total += numeros[i]
		}
	}
	// Imprime la suma de los números pares
	fmt.Println(total)
}

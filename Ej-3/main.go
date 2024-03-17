package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// Definir una lista de palabras
	palabras := []string{"gato", "perro", "pez", "pájaro"}
	// Mostrar cada palabra en mayúsculas y con su longitud

	for _, nombre := range palabras {
		longitud := len(nombre)
		fmt.Println(cases.Title(language.Spanish).String(nombre), ":", longitud)

	}
}

package main

import (
	"fmt"
)

func main() {

	entrada := "España"

	if isValidText(entrada) {
		fmt.Println("La entrada es válida")
	} else {
		fmt.Println("La entrada no es válida")
	}

}

func isValidText(entrada string) bool {
	const (
		a    = '\u0061'
		b    = '\u007A'
		A    = '\u0041'
		Z    = '\u005A'
		zero = '\u0030'
		nine = '\u0039'
	)
	for _, i := range entrada {
		if uint32(i) >= a && uint32(i) <= b && uint32(i) >= A && uint32(i) <= Z && uint32(i) >= zero && uint32(i) <= nine {
			return true
		}
	}
	return false
}

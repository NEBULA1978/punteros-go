package main

import "fmt"

func main() {
	// Declaración de una variable de tipo entero con un valor inicial de 4
	var v int = 4

	// Imprime el valor de la variable v en la cadena formateada
	fmt.Printf("El valor de v es %d\n", v)

	// Imprime la dirección de memoria de la variable v en la cadena formateada
	fmt.Printf("La posicion de memoria de v es %v\n", &v)
}

// Básicamente, el programa declara una variable de tipo entero llamada v y le asigna el valor 4. Luego, imprime el valor de v y su dirección de memoria usando las funciones Printf y & en la cadena formateada.

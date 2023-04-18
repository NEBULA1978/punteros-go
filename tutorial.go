package main

import "fmt"

func main() {
	// Declaración de una variable de tipo entero con un valor inicial de 4
	var v int = 4
	var p *int

	p = &v
	// El valor de la posicion de memoria al que apunta el puntero es 8
	*p = 8

	// Imprime el valor de la variable v en la cadena formateada
	fmt.Printf("El valor de v es %d \n", v)

	// Imprime la dirección de memoria de la variable v en la cadena formateada
	fmt.Printf("La posicion de memoria de v es %v \n", &v)

	// Imprime el valor que hay alojado en la posicion de memoria a la que apunta
	fmt.Printf("El valor de p es %d \n", *p)

	// Imprime la posicion de memoria de p
	// Los punteros simepre tienen una posicion de memoria no ponemos &
	fmt.Printf("La posicion de memoria de p es %v \n", p)
}

// ERROR1:
// RESULTADO POR CONSOLA:

// <▪> ~/V/G/punteros-go on master ✗ go run tutorial.go
// El valor de v es 4
// La posicion de memoria de v es 0xc00001a0f0
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4828fe]

// goroutine 1 [running]:
// main.main()
//         /home/next/Vídeos/GO/punteros-go/tutorial.go:17 +0xde
// exit status 2
// <▪> ~/V/G/punteros-go on master ✗

// & -> Para indicar la posicion de memoria
// // * -> Para indicar el valor

// Básicamente, el programa declara una variable de tipo entero llamada v y le asigna el valor 4. Luego, imprime el valor de v y su dirección de memoria usando las funciones Printf y & en la cadena formateada.

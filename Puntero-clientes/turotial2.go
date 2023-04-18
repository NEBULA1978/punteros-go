package main

import "fmt"

type Cliente struct {
	nombre   string
	telefono string
	next     *Cliente
}

var primerCliente *Cliente

func main() {
	AddCliente("Pepe", "444 444 444")
	AddCliente("Juan", "444 444 445")
	AddCliente("Luis", "444 444 446")
	AddCliente("Pedro", "444 444 446")
	ListaClientes()

	var cliente *Cliente = GetCliente(2)
	(*cliente).nombre = "Albert"

	fmt.Printf("\nClienes Update\n")
	ListaClientes()
}
func AddCliente(nombre string, telefono string) {
	var cliente = Cliente{
		nombre:   nombre,
		telefono: telefono,
		next:     nil,
	}
	if primerCliente == nil {
		primerCliente = &cliente

		return
	}
	var nextCliente *Cliente = primerCliente
	for nextCliente.next != nil {
		nextCliente = nextCliente.next
	}
	nextCliente.next = &cliente
}

func ListaClientes() {
	if primerCliente == nil {
		return
	}
	var nextCliente *Cliente = primerCliente
	for nextCliente.next != nil {
		fmt.Printf("Nombre: %s Teléfono: %s \n", nextCliente.nombre, nextCliente.telefono)
		nextCliente = nextCliente.next
	}
	fmt.Printf("Nombre: %s Teléfono: %s \n", nextCliente.nombre, nextCliente.telefono)
}

func GetCliente(pos int) *Cliente {
	if primerCliente == nil {
		return nil
	}

	var nextCliente *Cliente = primerCliente
	i := 0
	for i < pos && nextCliente.next != nil {
		nextCliente = nextCliente.next
		i++
	}
	if i != pos {
		return nil
	}
	return nextCliente
}

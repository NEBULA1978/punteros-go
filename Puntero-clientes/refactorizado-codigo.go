package main

import "fmt"

type Cliente struct {
	nombre   string
	telefono string
	next     *Cliente
}

var primerCliente *Cliente

func main() {
	fmt.Printf("\nClientes Agregados\n")
	AddCliente("Pepe", "444 444 444")
	AddCliente("Juan", "444 444 445")
	AddCliente("Luis", "444 444 446")
	AddCliente("Pedro", "444 444 446")
	ListarClientes()

	cliente := GetCliente(2)
	cliente.nombre = "Albert"

	fmt.Printf("\nClientes Update\n")
	ListarClientes()

	EliminarCliente(2)
	fmt.Printf("\nClientes Delete\n")
	ListarClientes()
}

func AddCliente(nombre string, telefono string) {
	cliente := &Cliente{
		nombre:   nombre,
		telefono: telefono,
		next:     nil,
	}
	if primerCliente == nil {
		primerCliente = cliente
		return
	}
	ultimoCliente := primerCliente
	for ultimoCliente.next != nil {
		ultimoCliente = ultimoCliente.next
	}
	ultimoCliente.next = cliente
}

func ListarClientes() {
	cliente := primerCliente
	for cliente != nil {
		fmt.Printf("Nombre: %s Teléfono: %s \n", cliente.nombre, cliente.telefono)
		cliente = cliente.next
	}
}

func GetCliente(pos int) *Cliente {
	cliente := primerCliente
	for i := 0; i < pos && cliente != nil; i++ {
		cliente = cliente.next
	}
	return cliente
}

func EliminarCliente(pos int) {
	if primerCliente == nil {
		return
	}
	if pos == 0 {
		primerCliente = primerCliente.next
		return
	}
	clienteAnterior := GetCliente(pos - 1)
	if clienteAnterior == nil || clienteAnterior.next == nil {
		return
	}
	clienteAnterior.next = clienteAnterior.next.next
}

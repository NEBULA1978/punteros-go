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

	DeleteCliente(2)
	fmt.Printf("\nClientes Delete\n")
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

func DeleteCliente(pos int) {
	if primerCliente == nil {
		return
	}
	if pos == 0 {
		if primerCliente.next == nil {
			primerCliente = nil
			return
		}
		primerCliente = primerCliente.next
		return
	}
	var cliente *Cliente = GetCliente(pos)
	var previo *Cliente = GetCliente(pos - 1)
	if cliente == nil || previo == nil {
		return
	}
	previo.next = cliente.next
}

func GetClienteByName(nombre string) *Cliente {
	if primerCliente == nil {
		return nil
	}
	var nextCliente *Cliente = primerCliente
	for nextCliente != nil {
		if nextCliente.nombre == nombre {
			return nextCliente
		}
		nextCliente = nextCliente.next
	}
	return nil
}

func UpdateCliente(nombre string, telefono string) {
	var cliente *Cliente = GetClienteByName(nombre)
	if cliente == nil {
		return
	}
	cliente.telefono = telefono
}

func CountClientes() int {
	if primerCliente == nil {
		return 0
	}
	var nextCliente *Cliente = primerCliente
	count := 1
	for nextCliente.next != nil {
		count++
		nextCliente = nextCliente.next
	}
	return count
}

// Este código es un ejemplo de una lista enlazada en Go. La estructura Cliente almacena información sobre un cliente, incluyendo su nombre y número de teléfono, y un puntero al siguiente cliente en la lista. La variable primerCliente es un puntero al primer elemento de la lista.

// Las funciones AddCliente, ListaClientes, GetCliente, DeleteCliente, GetClienteByName, UpdateCliente y CountClientes realizan diversas operaciones sobre la lista enlazada. AddCliente agrega un cliente a la lista, ListaClientes imprime todos los clientes en la lista, GetCliente devuelve un puntero al cliente en la posición especificada, DeleteCliente elimina el cliente en la posición especificada, GetClienteByName devuelve un puntero al cliente con el nombre especificado, UpdateCliente actualiza el teléfono de un cliente con el nombre especificado y CountClientes devuelve la cantidad de clientes en la lista.

// En la función main, se agregaron 4 clientes a la lista y se imprimieron. Luego se actualizó el nombre del segundo cliente y se eliminó el segundo cliente de la lista.

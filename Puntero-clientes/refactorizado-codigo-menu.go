package main

import (
	"fmt"
)

type Cliente struct {
	nombre    string
	telefono  string
	siguiente *Cliente
}

var primerCliente *Cliente

func main() {
	for {
		fmt.Println("Menú:")
		fmt.Println("1. Agregar cliente")
		fmt.Println("2. Actualizar cliente")
		fmt.Println("3. Eliminar cliente")
		fmt.Println("4. Listar clientes")
		fmt.Println("5. Salir")

		var opcion int
		fmt.Printf("Introduce una opción: ")
		fmt.Scanf("%d", &opcion)

		switch opcion {
		case 1:
			var nombre, telefono string
			fmt.Printf("Introduce el nombre del cliente: ")
			fmt.Scanf("%s", &nombre)
			fmt.Printf("Introduce el teléfono del cliente: ")
			fmt.Scanf("%s", &telefono)
			AgregarCliente(nombre, telefono)
		case 2:
			var pos int
			var nuevoNombre string
			fmt.Printf("Introduce la posición del cliente a actualizar: ")
			fmt.Scanf("%d", &pos)
			fmt.Printf("Introduce el nuevo nombre del cliente: ")
			fmt.Scanf("%s", &nuevoNombre)
			cliente := ObtenerCliente(pos)
			if cliente != nil {
				cliente.nombre = nuevoNombre
			} else {
				fmt.Println("Cliente no encontrado")
			}
		case 3:
			var pos int
			fmt.Printf("Introduce la posición del cliente a eliminar: ")
			fmt.Scanf("%d", &pos)
			EliminarCliente(pos)
		case 4:
			ListarClientes()
		case 5:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}

func AgregarCliente(nombre string, telefono string) {
	cliente := &Cliente{
		nombre:    nombre,
		telefono:  telefono,
		siguiente: nil,
	}
	if primerCliente == nil {
		primerCliente = cliente
		return
	}
	ultimoCliente := primerCliente
	for ultimoCliente.siguiente != nil {
		ultimoCliente = ultimoCliente.siguiente
	}
	ultimoCliente.siguiente = cliente
}

func ListarClientes() {
	cliente := primerCliente
	for cliente != nil {
		fmt.Printf("Nombre: %s Teléfono: %s \n", cliente.nombre, cliente.telefono)
		cliente = cliente.siguiente
	}
}

func ObtenerCliente(pos int) *Cliente {
	cliente := primerCliente
	for i := 0; i < pos && cliente != nil; i++ {
		cliente = cliente.siguiente
	}
	return cliente
}

func EliminarCliente(pos int) {
	if primerCliente == nil {
		return
	}
	if pos == 0 {
		primerCliente = primerCliente.siguiente
		return
	}
	clienteAnterior := ObtenerCliente(pos - 1)
	if clienteAnterior == nil || clienteAnterior.siguiente == nil {
		return
	}
	clienteAnterior.siguiente = clienteAnterior.siguiente.siguiente
}

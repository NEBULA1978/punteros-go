package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ...

func guardarClientesEnArchivo(nombreArchivo string) error {
	clientes := []Cliente{}
	cliente := primerCliente
	for cliente != nil {
		clientes = append(clientes, *cliente)
		cliente = cliente.siguiente
	}

	bytes, err := json.MarshalIndent(clientes, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(nombreArchivo, bytes, 0644)
}

func cargarClientesDesdeArchivo(nombreArchivo string) error {
	bytes, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		return err
	}

	clientes := []Cliente{}
	if err := json.Unmarshal(bytes, &clientes); err != nil {
		return err
	}

	primerCliente = nil
	for _, cliente := range clientes {
		AgregarCliente(cliente.nombre, cliente.telefono)
	}

	return nil
}
func main() {
	for {
		// ...
		fmt.Println("6. Guardar clientes en archivo")
		fmt.Println("7. Cargar clientes desde archivo")
		// ...

		switch opcion {
		// ...
		case 6:
			var nombreArchivo string
			fmt.Printf("Introduce el nombre del archivo para guardar los clientes: ")
			fmt.Scanf("%s", &nombreArchivo)
			if err := guardarClientesEnArchivo(nombreArchivo); err != nil {
				fmt.Printf("Error al guardar clientes: %v\n", err)
			} else {
				fmt.Println("Clientes guardados con éxito")
			}
		case 7:
			var nombreArchivo string
			fmt.Printf("Introduce el nombre del archivo para cargar los clientes: ")
			fmt.Scanf("%s", &nombreArchivo)
			if err := cargarClientesDesdeArchivo(nombreArchivo); err != nil {
				fmt.Printf("Error al cargar clientes: %v\n", err)
			} else {
				fmt.Println("Clientes cargados con éxito")
			}
			// ...
		}
	}
}

// Con estas funciones y opciones de menú, podrás guardar y cargar los clientes en un archivo, lo que te permitirá conservar la información de los clientes entre diferentes ejecuciones del programa.

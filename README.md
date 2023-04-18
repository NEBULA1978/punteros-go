# punteros-go

Ejemplo de lista enlazada simple en Go

Este código en Go es un ejemplo de cómo trabajar con una lista enlazada simple. La lista enlazada
está compuesta por una serie de nodos, donde cada nodo contiene un dato y una referencia al siguiente nodo de la lista.

En este ejemplo, cada nodo representa un cliente y tiene dos campos: el nombre y el número de teléfono.
Además, cada nodo tiene una referencia al siguiente nodo de la lista.
Funciones incluidas

El programa incluye las siguientes funciones para trabajar con la lista de clientes:

    AddCliente(nombre string, telefono string): Añade un nuevo cliente a la lista.

    ListaClientes(): Muestra la lista de clientes por pantalla.

    *GetCliente(pos int) Cliente: Devuelve el cliente en la posición indicada.

    DeleteCliente(pos int): Elimina el cliente en la posición indicada de la lista.

Uso del programa

El programa comienza definiendo la estructura de Cliente y una variable global llamada primerCliente 
que apunta al primer nodo de la lista (o a nil si la lista está vacía).

En la función main, se agregan cuatro clientes a la lista y se muestran por pantalla. Luego, se obtiene
el cliente en la posición 2, se actualiza su nombre y se muestra la lista actualizada. Finalmente,
se elimina el cliente en la posición 2 y se muestra la lista de clientes restante.

En general, la lista enlazada es una estructura de datos muy útil en programación y se utiliza 
en una amplia variedad de aplicaciones. Este ejemplo muestra cómo se puede trabajar con una lista
enlazada simple en Go y puede servir como punto de partida para crear aplicaciones más complejas 
que requieren el uso de esta estructura de datos.

package lista

type Lista[T any] interface {

	//EstaVacia devuelve verdadero si la lista esta vacia
	EstaVacia() bool

	//InsertarPrimero inserta un elemento al principio de la lista
	InsertarPrimero(T)

	//InsertarUltimo inserta un elemento al final de la lista
	InsertarUltimo(T)

	//BorrarPrimero elimina el primer elemento de la lista y devuelve ese valor. Si la lista esta vacia
	//Entra en panico con un mensaje "la lista esta vacia"
	BorrarPrimero() T

	//VerPrimero devuelve el primer elemento de la lista. Si la lista esta vacia
	//Entra en panico con un mensaje "la lista esta vacia"
	VerPrimero() T

	//VerUltimo devuelve el ultimo elemento de la lista. Si la lista esta vacia
	//Entra en panico con un mensaje "la lista esta vacia"
	VerUltimo() T

	//Largo devuelve el largo de la lista.
	Largo() int

	//Iterar itera de manera interna a la lista
	Iterar(visitar func(T) bool)

	//Iterador devuelve un iterador externo
	Iterador() IteradorLista[T]
}
package lista


type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int

}

func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato, siguiente: nil}

}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elementoParaInsertar T){

	nuevo_nodo := crearNodo(elementoParaInsertar)

	if lista.EstaVacia() {
		lista.primero = nuevo_nodo
		lista.ultimo = nuevo_nodo

	} else {
		nuevo_nodo.siguiente = lista.primero
		lista.primero = nuevo_nodo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elementoParaInsertar T){

	nuevo_nodo := crearNodo(elementoParaInsertar)

	if lista.EstaVacia(){
		lista.primero = nuevo_nodo
		lista.ultimo = nuevo_nodo
	} else {
		lista.ultimo.siguiente = nuevo_nodo
		lista.ultimo = nuevo_nodo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T{

	if lista.EstaVacia(){
		panic("la lista esta vacia")
	}

	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente

	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo --

	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T{
	
	if lista.EstaVacia(){
		panic("la lista esta vacia")
	}

	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T{
	if lista.EstaVacia(){
		panic("la lista esta vacia")
	}

	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int{
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero

	for actual != nil {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{actual: lista.primero}
}

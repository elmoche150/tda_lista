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

type iteradorLista[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T]{
		actual:   lista.primero,
		anterior: nil,
		lista:    lista,
	}
}

func (iter *iteradorLista[T]) HayAlgoMas() bool {
	return iter.actual != nil
}

func (iter *iteradorLista[T]) VerActual() T {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iteradorLista[T]) Avanzar() {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorLista[T]) Insertar(valor T) {
	nuevo := &nodoLista[T]{dato: valor}

	// insertar al inicio
	if iter.anterior == nil {
		nuevo.siguiente = iter.lista.primero
		iter.lista.primero = nuevo

	} else {
		// esta en el medio
		nuevo.siguiente = iter.actual
		iter.anterior.siguiente = nuevo
	}

	// si esta al final
	if iter.actual == nil {
		iter.lista.ultimo = nuevo
	}

	iter.actual = nuevo
	iter.lista.largo++
}

func (iter *iteradorLista[T]) Borrar() T {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}

	borrado := iter.actual

	//borrar primero
	if iter.anterior == nil {
		iter.lista.primero = borrado.siguiente
	} else {
		iter.anterior.siguiente = borrado.siguiente
	}

	// borrar ultimo
	if borrado.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}

	iter.actual = borrado.siguiente
	iter.lista.largo--

	return borrado.dato
}

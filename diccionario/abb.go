package diccionario

import "tdas/pila"

const (
	CANTIDAD_INICIAL  = 0
	COMPARACION_MENOR = 0
	COMPARACION_MAYOR = 0
	COMPARACION_IGUAL = 0
)

type funcCmp[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type iterAbb[K comparable, V any] struct {
	pila  pila.Pila[*nodoAbb[K, V]]
	abb   *abb[K, V]
	desde *K
	hasta *K
}

func (a *abb[K, V]) guardarNodo(actual *nodoAbb[K, V], clave K, dato V) *nodoAbb[K, V] {
	if actual == nil {
		a.cantidad++
		return &nodoAbb[K, V]{
			clave: clave,
			dato:  dato,
		}
	}
	comparacion := a.cmp(clave, actual.clave)

	if comparacion < COMPARACION_MENOR {

		actual.izquierdo = a.guardarNodo(actual.izquierdo, clave, dato)

	} else if comparacion > COMPARACION_MAYOR {

		actual.derecho = a.guardarNodo(actual.derecho, clave, dato)

	} else {

		actual.dato = dato
	}

	return actual
}

func (a *abb[K, V]) perteneceNodo(actual *nodoAbb[K, V], clave K) bool {
	if actual == nil {
		return false
	}

	comparacion := a.cmp(clave, actual.clave)

	if comparacion < COMPARACION_MENOR {
		return a.perteneceNodo(actual.izquierdo, clave)
	}

	if comparacion > COMPARACION_MAYOR {
		return a.perteneceNodo(actual.derecho, clave)
	}

	return true

}

func (a *abb[K, V]) obtenerNodo(actual *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	if actual == nil {
		return nil
	}

	comparacion := a.cmp(clave, actual.clave)

	if comparacion < COMPARACION_MENOR {
		return a.obtenerNodo(actual.izquierdo, clave)
	}

	if comparacion > COMPARACION_MAYOR {
		return a.obtenerNodo(actual.derecho, clave)
	}

	return actual
}

func buscarMinimo[K comparable, V any](actual *nodoAbb[K, V]) *nodoAbb[K, V] {
	if actual.izquierdo == nil {
		return actual
	}

	return buscarMinimo(actual.izquierdo)
}

// borrarNodo ahora devuelve (nuevoNodo, valor, encontrado)
// borrarNodo ahora devuelve (nuevoNodo, valor, encontrado)
func (a *abb[K, V]) borrarNodo(actual *nodoAbb[K, V], clave K) (*nodoAbb[K, V], V, bool) {
	if actual == nil {
		var cero V
		return nil, cero, false // No lo encontré, aviso con el 'false'
	}

	comparacion := a.cmp(clave, actual.clave)

	if comparacion < COMPARACION_MENOR {
		izq, dato, ok := a.borrarNodo(actual.izquierdo, clave)
		actual.izquierdo = izq
		return actual, dato, ok
	}

	if comparacion > COMPARACION_MAYOR {
		der, dato, ok := a.borrarNodo(actual.derecho, clave)
		actual.derecho = der
		return actual, dato, ok
	}

	// ¡Lo encontramos! (comparacion == 0)
	datoBorrado := actual.dato

	// Caso 1: Hoja o un solo hijo
	if actual.izquierdo == nil {
		return actual.derecho, datoBorrado, true
	}
	if actual.derecho == nil {
		return actual.izquierdo, datoBorrado, true
	}

	// Caso 2: Dos hijos (buscamos reemplazo con el sucesor)
	sucesor := buscarMinimo(actual.derecho)
	actual.clave = sucesor.clave
	actual.dato = sucesor.dato

	// Borramos el sucesor en el subárbol derecho
	// Nota: acá no necesitamos el booleano porque sabemos que el sucesor existe
	der, _, _ := a.borrarNodo(actual.derecho, sucesor.clave)
	actual.derecho = der

	return actual, datoBorrado, true
}

func (a *abb[K, V]) iterarRangoNodo(actual *nodoAbb[K, V], desde *K, hasta *K, visitar func(K, V) bool) bool {

	if actual == nil {
		return true
	}

	if desde == nil || a.cmp(actual.clave, *desde) > 0 {
		if !a.iterarRangoNodo(actual.izquierdo, desde, hasta, visitar) {
			return false
		}
	}

	if (desde == nil || a.cmp(actual.clave, *desde) >= 0) &&
		(hasta == nil || a.cmp(actual.clave, *hasta) <= 0) {

		if !visitar(actual.clave, actual.dato) {
			return false
		}
	}

	if hasta == nil || a.cmp(actual.clave, *hasta) < 0 {
		if !a.iterarRangoNodo(actual.derecho, desde, hasta, visitar) {
			return false
		}
	}

	return true
}

func (iter *iterAbb[K, V]) apilarRamaIzquierdaRango(actual *nodoAbb[K, V]) {

	for actual != nil {

		if iter.desde != nil &&
			iter.abb.cmp(actual.clave, *iter.desde) < 0 {

			actual = actual.derecho

		} else {

			iter.pila.Apilar(actual)
			actual = actual.izquierdo
		}
	}
}

func (iter *iterAbb[K, V]) HayAlgoMas() bool {

	if iter.pila.EstaVacia() {
		return false
	}

	if iter.hasta != nil {

		actual := iter.pila.VerTope()

		if iter.abb.cmp(actual.clave, *iter.hasta) > 0 {
			return false
		}
	}

	return true
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {

	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	actual := iter.pila.VerTope()

	return actual.clave, actual.dato
}

func (iter *iterAbb[K, V]) Avanzar() {

	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	actual := iter.pila.Desapilar()

	if actual.derecho != nil {
		iter.apilarRamaIzquierdaRango(actual.derecho)
	}
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: CANTIDAD_INICIAL, cmp: funcion_cmp}
}

func (a *abb[K, V]) Guardar(clave K, dato V) {
	a.raiz = a.guardarNodo(a.raiz, clave, dato)
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	return a.perteneceNodo(a.raiz, clave)
}

func (a *abb[K, V]) Obtener(clave K) V {
	elemento := a.obtenerNodo(a.raiz, clave)
	if elemento == nil {
		panic("La clave no pertenece al diccionario")
	}
	return elemento.dato
}

func (a *abb[K, V]) Borrar(clave K) V {

	nuevaRaiz, datoBorrado, encontrado := a.borrarNodo(a.raiz, clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}

	a.raiz = nuevaRaiz
	a.cantidad--
	return datoBorrado
}

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

	if desde != nil && hasta != nil &&
		a.cmp(*desde, *hasta) > 0 {
		return
	}

	a.iterarRangoNodo(a.raiz, desde, hasta, visitar)
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

	iter := &iterAbb[K, V]{
		pila:  pila.CrearPilaDinamica[*nodoAbb[K, V]](),
		abb:   a,
		desde: desde,
		hasta: hasta,
	}

	if desde != nil && hasta != nil &&
		a.cmp(*desde, *hasta) > 0 {

		return iter
	}

	iter.apilarRamaIzquierdaRango(a.raiz)

	return iter
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

func (a *abb[K, V]) Iterar(visitar func(K, V) bool) {
	a.IterarRango(nil, nil, visitar)
}

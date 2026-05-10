package diccionario

const (
	CANTIDAD_INICIAL   = 0
	COMPARACION_MENOR  = 0
	COMPARACION_MAYOR  = 0
	COMPPARACION_IGUAL = 0
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
	}

	if comparacion > COMPARACION_MAYOR {
		actual.derecho = a.guardarNodo(actual.derecho, clave, dato)
	}

	if comparacion == COMPPARACION_IGUAL {
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

func (a *abb[K, V]) borrarNodo(actual *nodoAbb[K, V], clave K) (*nodoAbb[K, V], V) {
	if actual == nil {
		var cero V
		return nil, cero
	}
	comparacion := a.cmp(clave, actual.clave)

	if comparacion < COMPARACION_MENOR {
		var dato V

		actual.izquierdo, dato = a.borrarNodo(actual.izquierdo, clave)
		return actual, dato
	}

	if comparacion > COMPARACION_MAYOR {
		var dato V

		actual.derecho, dato = a.borrarNodo(actual.derecho, clave)
		return actual, dato
	}

	datoBorrado := actual.dato

	if actual.izquierdo == nil && actual.derecho == nil {
		return nil, datoBorrado
	}

	if actual.izquierdo == nil {
		return actual.derecho, datoBorrado
	}

	if actual.derecho == nil {
		return actual.izquierdo, datoBorrado
	}

	sucesor := buscarMinimo(actual.derecho)

	actual.clave = sucesor.clave
	actual.dato = sucesor.dato

	actual.derecho, _ = a.borrarNodo(actual.derecho, sucesor.clave)

	return actual, datoBorrado
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

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: CANTIDAD_INICIAL, cmp: funcionCmp}
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
		print("no existe la clave")
		panic("no existe la clave")
	}
	return elemento.dato
}

func (a *abb[K, V]) Borrar(clave K) V {
	if !a.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}

	var datoBorrado V

	a.raiz, datoBorrado = a.borrarNodo(a.raiz, clave)

	a.cantidad--
	return datoBorrado
}

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	a.iterarRangoNodo(a.raiz, desde, hasta, visitar)
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iterAbb[K, V]{
		pila:  pila.CrearPilaDinamica[*nodoAbb[K, V]](),
		abb:   a,
		desde: desde,
		hasta: hasta,
	}

	iter.apilarRamaIzquierdaRango(a.raiz)

	return iter
}

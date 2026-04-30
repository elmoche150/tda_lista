package diccionario

import (
	"fmt"

	"tdas/lista"
)

const (
	CAPACIDAD_INICIAL           = 10
	CANTIDAD_INICIAL            = 0
	PARAMETRO_PARA_AGRANDAR     = 2
	CANTIDAD_PARA_REDIMENSIONAR = 2
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []lista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iterHash[K comparable, V any] struct {
	hash         *hashAbierto[K, V]
	iterLista    lista.IteradorLista[parClaveValor[K, V]]
	indiceActual int
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func funcionHashing(bytes []byte) int {
	hash := 0

	for _, b := range bytes {
		hash = hash*31 + int(b)
	}

	if hash < 0 {
		hash = -hash
	}

	return hash
}

func (hash *hashAbierto[K, V]) indice(clave K) int {
	return funcionHashing(convertirABytes(clave)) % hash.tam
}

func crearTabla[K comparable, V any](tam int) []lista.Lista[parClaveValor[K, V]] {
	tabla := make([]lista.Lista[parClaveValor[K, V]], tam)

	for i := range tabla {
		tabla[i] = lista.CrearListaEnlazada[parClaveValor[K, V]]()
	}

	return tabla
}
func (hash *hashAbierto[K, V]) chequearRedimension() {
	factorCarga := float64(hash.cantidad) / float64(hash.tam)

	if factorCarga > PARAMETRO_PARA_AGRANDAR {
		hash.redimensionar(hash.tam * CANTIDAD_PARA_REDIMENSIONAR)
	} else if factorCarga < 0.25 && hash.tam > CAPACIDAD_INICIAL {
		nuevoTam := hash.tam / CANTIDAD_PARA_REDIMENSIONAR
		if nuevoTam < CAPACIDAD_INICIAL {
			nuevoTam = CAPACIDAD_INICIAL
		}
		hash.redimensionar(nuevoTam)
	}
}

func (hash *hashAbierto[K, V]) redimensionar(nuevoTam int) {
	tablaVieja := hash.tabla

	hash.tabla = crearTabla[K, V](nuevoTam)
	hash.tam = nuevoTam
	hash.cantidad = CANTIDAD_INICIAL

	for _, l := range tablaVieja {
		l.Iterar(func(par parClaveValor[K, V]) bool {
			hash.Guardar(par.clave, par.dato)
			return true
		})
	}
}

func (hash *hashAbierto[K, V]) borrarDeLista(indice int, clave K) {
	nueva := lista.CrearListaEnlazada[parClaveValor[K, V]]()
	hash.tabla[indice].Iterar(func(par parClaveValor[K, V]) bool {
		if par.clave != clave {
			nueva.InsertarUltimo(par)
		}
		return true
	})
	hash.tabla[indice] = nueva
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {

	tabla := crearTabla[K, V](CAPACIDAD_INICIAL)

	return &hashAbierto[K, V]{tabla: tabla, tam: CAPACIDAD_INICIAL, cantidad: CANTIDAD_INICIAL}
}

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	indice := hash.indice(clave)
	encontrado := false

	hash.tabla[indice].Iterar(func(par parClaveValor[K, V]) bool {
		if par.clave == clave {
			encontrado = true
			return false
		}
		return true
	})

	if encontrado {
		hash.borrarDeLista(indice, clave)
	} else {
		hash.cantidad++
	}

	hash.tabla[indice].InsertarUltimo(parClaveValor[K, V]{clave, dato})
	hash.chequearRedimension()
}
func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	indice := hash.indice(clave)

	var dato V
	encontrado := false

	hash.tabla[indice].Iterar(func(par parClaveValor[K, V]) bool {
		if par.clave == clave {
			dato = par.dato
			encontrado = true
			return false
		}
		return true
	})

	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}

	return dato
}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	indice := hash.indice(clave)

	encontrado := false

	hash.tabla[indice].Iterar(func(par parClaveValor[K, V]) bool {
		if par.clave == clave {
			encontrado = true
			return false
		}
		return true
	})

	return encontrado
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	indice := hash.indice(clave)
	var dato V
	encontrado := false

	hash.tabla[indice].Iterar(func(par parClaveValor[K, V]) bool {
		if par.clave == clave {
			dato = par.dato
			encontrado = true
			return false
		}
		return true
	})

	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}

	hash.borrarDeLista(indice, clave)
	hash.cantidad--

	hash.chequearRedimension()

	return dato
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (hash *hashAbierto[K, V]) Iterar(f func(clave K, dato V) bool) {
	for _, l := range hash.tabla {
		salir := false
		l.Iterar(func(par parClaveValor[K, V]) bool {
			if !f(par.clave, par.dato) {
				salir = true
				return false
			}
			return true
		})
		if salir {
			break
		}
	}
}

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterHash[K, V]{hash: hash, indiceActual: 0}

	for iter.indiceActual < iter.hash.tam && iter.hash.tabla[iter.indiceActual].EstaVacia() {
		iter.indiceActual++
	}

	if iter.indiceActual < iter.hash.tam {
		iter.iterLista = iter.hash.tabla[iter.indiceActual].Iterador()
	}

	return iter
}

func (iter *iterHash[K, V]) HayAlgoMas() bool {
	return iter.iterLista != nil && iter.iterLista.HayAlgoMas()
}

func (iter *iterHash[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	par := iter.iterLista.VerActual()
	return par.clave, par.dato
}

func (iter *iterHash[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	iter.iterLista.Avanzar()

	if !iter.iterLista.HayAlgoMas() {
		iter.indiceActual++
		for iter.indiceActual < iter.hash.tam && iter.hash.tabla[iter.indiceActual].EstaVacia() {
			iter.indiceActual++
		}

		if iter.indiceActual < iter.hash.tam {
			iter.iterLista = iter.hash.tabla[iter.indiceActual].Iterador()
		} else {
			iter.iterLista = nil
		}
	}
}

package lista_test

import (
	"testing"

	lista "../tda_lista"

	"github.com/stretchr/testify/require"
)

func crearListasParaTestear() (lista.Lista[int], lista.Lista[string], lista.Lista[float64]) {
	lista_enteros := lista.CrearListaEnlazada[int]()
	lista_strings := lista.CrearListaEnlazada[string]()
	lista_float64 := lista.CrearListaEnlazada[float64]()

	return lista_enteros, lista_strings, lista_float64
}

func TestListaVacia(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	require.True(t, lista_enteros.EstaVacia())
	require.True(t, lista_strings.EstaVacia())
	require.True(t, lista_float64.EstaVacia())
}

func TestInsertarPrimero(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	require.False(t, lista_enteros.EstaVacia())

	lista_strings.InsertarPrimero("martu")
	require.False(t, lista_strings.EstaVacia())

	lista_float64.InsertarPrimero(10.2)
	require.False(t, lista_enteros.EstaVacia())
}

func TestInsertarUltimo(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarUltimo(10)
	require.False(t, lista_enteros.EstaVacia())

	lista_strings.InsertarUltimo("martu")
	require.False(t, lista_strings.EstaVacia())

	lista_float64.InsertarUltimo(10.2)
	require.False(t, lista_float64.EstaVacia())
}

func TestVerPrimero(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.InsertarUltimo(110)
	require.Equal(t, 10, lista_enteros.VerPrimero())

	lista_strings.InsertarPrimero("martu")
	lista_strings.InsertarUltimo("paloma")
	require.Equal(t, "martu", lista_strings.VerPrimero())

	lista_float64.InsertarPrimero(1.23)
	lista_float64.InsertarUltimo(1.55)
	require.Equal(t, 1.23, lista_float64.VerPrimero())

}

func TestVerUltimo(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarUltimo(10)
	lista_enteros.InsertarPrimero(110)
	require.Equal(t, 10, lista_enteros.VerUltimo())

	lista_strings.InsertarUltimo("martu")
	lista_strings.InsertarPrimero("paloma")
	require.Equal(t, "martu", lista_strings.VerUltimo())

	lista_float64.InsertarUltimo(1.23)
	lista_float64.InsertarPrimero(1.55)
	require.Equal(t, 1.23, lista_float64.VerUltimo())
}

func TestVerPrimeroEnUnaListaVacia(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.BorrarPrimero()
	require.True(t, lista_enteros.EstaVacia())
	require.Panics(t, func() { lista_enteros.VerPrimero() })

	lista_strings.InsertarPrimero("moche")
	lista_strings.BorrarPrimero()
	require.True(t, lista_strings.EstaVacia())
	require.Panics(t, func() { lista_strings.VerPrimero() })

	lista_float64.InsertarPrimero(10.2)
	lista_float64.BorrarPrimero()
	require.True(t, lista_float64.EstaVacia())
	require.Panics(t, func() { lista_float64.VerPrimero() })
}

func TestVerUltimoEnUnaListaVacia(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.BorrarPrimero()
	require.True(t, lista_enteros.EstaVacia())
	require.Panics(t, func() { lista_enteros.VerUltimo() })

	lista_strings.InsertarPrimero("moche")
	lista_strings.BorrarPrimero()
	require.True(t, lista_strings.EstaVacia())
	require.Panics(t, func() { lista_strings.VerUltimo() })

	lista_float64.InsertarPrimero(10.2)
	lista_float64.BorrarPrimero()
	require.True(t, lista_float64.EstaVacia())
	require.Panics(t, func() { lista_float64.VerUltimo() })
}

func TestBorrarPrimero(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.BorrarPrimero()
	require.True(t, lista_enteros.EstaVacia())

	lista_strings.InsertarPrimero("martu")
	lista_strings.BorrarPrimero()
	require.True(t, lista_strings.EstaVacia())

	lista_float64.InsertarPrimero(10.2)
	lista_float64.BorrarPrimero()
	require.True(t, lista_float64.EstaVacia())
}

func TestBorrarPrimeroConvierteAlSegundoEnPrimero(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.InsertarUltimo(100)
	lista_enteros.InsertarUltimo(10000)
	lista_enteros.BorrarPrimero()
	require.Equal(t, 100, lista_enteros.VerPrimero())

	lista_strings.InsertarPrimero("coscu")
	lista_strings.InsertarUltimo("nachoide")
	lista_strings.InsertarUltimo("nanoide")
	lista_strings.BorrarPrimero()
	require.Equal(t, "nachoide", lista_strings.VerPrimero())

	lista_float64.InsertarPrimero(1.23)
	lista_float64.InsertarUltimo(1.2343)
	lista_float64.InsertarUltimo(2.34)
	lista_float64.BorrarPrimero()
	require.Equal(t, 1.2343, lista_float64.VerPrimero())
}

func TestLargo(t *testing.T) {
	lista_enteros, lista_strings, lista_float64 := crearListasParaTestear()

	lista_enteros.InsertarPrimero(10)
	lista_enteros.InsertarUltimo(100)
	lista_enteros.InsertarUltimo(10000)
	require.Equal(t, 3, lista_enteros.Largo())

	lista_strings.InsertarPrimero("coscu")
	lista_strings.InsertarUltimo("nachoide")
	lista_strings.InsertarUltimo("nanoide")
	require.Equal(t, 3, lista_strings.Largo())

	lista_float64.InsertarPrimero(1.23)
	lista_float64.InsertarUltimo(1.2343)
	lista_float64.InsertarUltimo(2.34)
	require.Equal(t, 3, lista_float64.Largo())
}

func TestVolumen(t *testing.T) {
	lista_enteros := lista.CrearListaEnlazada[int]()

	for i := 0; i < 10000; i++ {
		lista_enteros.InsertarPrimero(i)
		require.Equal(t, i, lista_enteros.VerPrimero())
	}

	for i := 0; i < 10000; i++ {
		valor := lista_enteros.BorrarPrimero()
		require.Equal(t, i, valor)
	}

	require.True(t, lista_enteros.EstaVacia())

}

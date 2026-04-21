package lista

type IteradorLista[T any] interface {
	VerActual() T
	HayAlgoMas() bool
	Avanzar()
	Insertar(T)
	Borrar() T
}

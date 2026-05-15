package cola_prioridad

type heap[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func (h *heap[T]) Cantidad() int {
	return h.cant
}

func (h *heap[T]) EstaVacia() bool {
	return h.cant == 0
}

func CrearHeap[T any](cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{
		datos: make([]T, 10),
		cmp:   cmp,
	}
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

func (h *heap[T]) Encolar(elem T) {
	h.datos[h.cant] = elem
	h.cant++

	//upheap()
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}

	max := h.datos[0]

	h.cant--
	h.datos[0] = h.datos[h.cant]
	//dowheap()

	return max
}

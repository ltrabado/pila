package pila

const _CAPACIDADINICIAL = 10

/* Definición del struct pila proporcionado por la cátedra. */
type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

/* Constructor del TAD Pila. */
func CrearPilaDinamica[T any]() Pila[T] {
	//const _CAPACIDADINICIAL = 10
    pila := new(pilaDinamica[T]) // se agrega el generic
    // hago lo que deba hacer
	pila.datos = make([]T, _CAPACIDADINICIAL)
	pila.cantidad = 0
    return pila
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func (pila *pilaDinamica[T]) EstaVacia() bool {
	if pila.cantidad == 0 {
		return true
	}
	return false
}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() == true {
		panic("La pila esta vacia") 
	}
	return pila.datos[pila.cantidad-1]
}

/* redimensionar ajusta la capacidad de la pila de la siguiente manera:
	1) Si cantidad es igual a la capacidad de la pila, se duplica su capacidad.
	2) Si cantidad es menor o igual a la cuarta parte de la capacidad de la pila, se reduce su capacidad a la mitad. */

func (pila *pilaDinamica[T]) redimensionar() {
	if cap(pila.datos) > _CAPACIDADINICIAL {
		if pila.cantidad == cap(pila.datos) {
			//Duplicar capacidad
			nuevoVector := make([]T, pila.cantidad*2)
			copy(nuevoVector, pila.datos)
			pila.datos = nuevoVector
		}
		if (pila.cantidad*4) <= cap(pila.datos) {
			//capacidad=capacidad/2
			nuevoVector := make([]T, pila.cantidad/2)
			copy(nuevoVector, pila.datos)
			pila.datos = nuevoVector
		}
	}
	return
}

// Apilar agrega un nuevo elemento a la pila.
func (pila *pilaDinamica[T]) Apilar(elem T) {
	if pila.cantidad == cap(pila.datos) {
		pila.redimensionar()
	}
	pila.datos[pila.cantidad] = elem
	pila.cantidad++
	return
}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() == true {
		panic("La pila esta vacia") 
	}
	elementoDesapilado := pila.datos[pila.cantidad-1]
	pila.cantidad--
	pila.redimensionar()
	return elementoDesapilado
}
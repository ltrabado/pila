package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

/* Constructor del TAD Pila. */

func CrearPilaDinamica[T any]() Pila[T] {
    pila := new(pilaDinamica[T]) // se agrega el generic
    // hago lo que deba hacer
    return pila
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func EstaVacia() bool {
	if cantidad == 0 {
		return true
	}
	return false
}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func VerTope() T {
	if EstaVacia() == true {
		panic("La pila esta vacia") 
	}
	return datos[cantidad - 1]
}

// Apilar agrega un nuevo elemento a la pila.
Apilar(T)

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func Desapilar() T {
	if EstaVacia() == true {
		panic("La pila esta vacia") 
	}
	elementoDesapilado := datos[cantidad - 1]
	cantidad--
	return elementoDesapilado
}
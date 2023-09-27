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
	return datos[cantidad-1]
}

/* redimensionar ajusta la capacidad de la pila de la siguiente manera:
	1) Si la pila se encuentra recien creada (pila vacia y capacidad igual a cero), se le asigna una capacidad inicial de 10.
	2) Si cantidad es igual a la capacidad de la pila, se duplica su capacidad.
	3) Si cantidad es menor o igual a la cuarta parte de la capacidad de la pila, se reduce su capacidad a la mitad. */
func redimensionar() T {
	if (EstaVacia() == (cap(datos)==0)) {
		//Asignar capacidad = 10
	}
	if cantidad == cap(datos) {
		//Duplicar capacidad
	}
	if (cantidad*4) <= cap(datos) {
		//capacidad=capacidad/2
	}
}
// Apilar agrega un nuevo elemento a la pila.
func Apilar(T) {
	if cantidad == cap(datos) {
		//redimensionar
	}
	datos[cantidad] = T
	cantidad++
}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func Desapilar() T {
	if EstaVacia() == true {
		panic("La pila esta vacia") 
	}
	elementoDesapilado := datos[cantidad-1]
	cantidad--
	//redimensionar
	return elementoDesapilado
}
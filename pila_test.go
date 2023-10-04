package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	// mas pruebas para este caso...
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestInvariantePila(t *testing.T) {
	TAM := 10
	pila := TDAPila.CrearPilaDinamica[int]()
	//Apilo enteros del 0 al 9
	for i := 0; i < TAM; i++ {
		pila.Apilar(i)
	}
	//Desapilo manteniendo el invariante de la pila
	for i := TAM-1; i >= 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}
}

func TestPilaFloat64Vacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestInvariantePilaFloat64(t *testing.T) {
	TAM := 10
	PI := 3.14
	pila := TDAPila.CrearPilaDinamica[float64]()
	//Apilo flotantes
	for i := 0; i < TAM; i++ {
		pila.Apilar(float64(i) * PI)
	}
	//Desapilo flotantes
	for i := TAM-1; i >= 0; i-- {
		require.EqualValues(t, float64(i) * PI, pila.Desapilar())
	}
}

func TestPilaBooleanaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[bool]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestInvariantePilaBooleana(t *testing.T) {
	TAM := 10
	pila := TDAPila.CrearPilaDinamica[bool]()
	//Apilo flotantes
	for i := 0; i < TAM; i++ {
		if i % 2 == 0 {
			pila.Apilar(true)
		} else {
			pila.Apilar(false)
		}
	}
	//Desapilo flotantes
	for i := TAM-1; i >= 0; i-- {
		if i % 2 == 0 {
			require.True(t, pila.Desapilar())
		} else {
			require.False(t, pila.Desapilar())
		}
	}
}

func TestPilaStringVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestInvariantePilaString(t *testing.T) {
	mensaje := []string{"Esta", "es", "una", "pila", "de", "cadenas"}
	mensajeInvertido := []string{"cadenas", "de", "pila", "una", "es", "Esta"}
	pila := TDAPila.CrearPilaDinamica[string]()
	//Apilo cadenas
	for i := range mensaje {
		pila.Apilar(mensaje[i])
	}
	//Desapilo cadenas
	for i := range mensajeInvertido {
		require.EqualValues(t, mensajeInvertido[i], pila.Desapilar())
	}
}
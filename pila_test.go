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
	for i := TAM - 1; i >= 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}
}

func TestVolumen(t *testing.T) {
	TAM := 10000
	pila := TDAPila.CrearPilaDinamica[int]()

	t.Log("Prueba de redimension creciente")
	//Apilo
	for i := 0; i < TAM; i++ {
		pila.Apilar(i)
	}

	t.Log("Prueba de redimension decreciente")
	//Desapilo
	for i := TAM - 1; i >= 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}
}

func TestDesapiloHastaVaciar(t *testing.T) {
	TAM := 15
	pila := TDAPila.CrearPilaDinamica[int]()
	//Apilo
	for i := 0; i < TAM; i++ {
		pila.Apilar(i)
	}
	//Desapilo
	for i := TAM - 1; i >= 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
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
	for i := TAM - 1; i >= 0; i-- {
		require.EqualValues(t, float64(i)*PI, pila.Desapilar())
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
		if i%2 == 0 {
			pila.Apilar(true)
		} else {
			pila.Apilar(false)
		}
	}
	//Desapilo flotantes
	for i := TAM - 1; i >= 0; i-- {
		if i%2 == 0 {
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

func TestPruebaBordeRedimension(t *testing.T) {
	TAM := 10000
	muestra := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	pila := TDAPila.CrearPilaDinamica[int]()

	t.Log("Apilo elementos y a continuacion los desapilo")
	//Apilo y desapilo
	for i := 0; i < TAM; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.Desapilar())
	}

	t.Log("Apilo y desapilo en las iteraciones pares")
	//Apilo la muestra y desapilo en las iteraciones pares
	for i := range muestra {
		pila.Apilar(muestra[i])
		if i%2 == 0 {
			require.EqualValues(t, i, pila.Desapilar())
		}
	}

	t.Log("Desapilo los elementos restantes")
	//Desapilo los elementos restantes
	require.EqualValues(t, 15, pila.Desapilar())
	require.EqualValues(t, 13, pila.Desapilar())
	require.EqualValues(t, 11, pila.Desapilar())
	require.EqualValues(t, 9, pila.Desapilar())
	require.EqualValues(t, 7, pila.Desapilar())
	require.EqualValues(t, 5, pila.Desapilar())
	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 1, pila.Desapilar())

	t.Log("Reviso que la pila este vacia")
	require.True(t, pila.EstaVacia())
}
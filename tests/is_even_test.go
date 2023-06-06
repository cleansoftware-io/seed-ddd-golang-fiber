package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEven(t *testing.T) {
	// Caso de prueba 1: Número par
	result := IsEven(4)
	assert.True(t, result, "4 debería ser un número par")

	// Caso de prueba 2: Número impar
	result = IsEven(7)
	assert.False(t, result, "7 no debería ser un número par")

	// Caso de prueba 3: Número negativo par
	result = IsEven(-10)
	assert.True(t, result, "-10 debería ser un número par")

	// Caso de prueba 4: Número negativo impar
	result = IsEven(-13)
	assert.False(t, result, "-13 no debería ser un número par")
}

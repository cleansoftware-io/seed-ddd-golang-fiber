package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEven(t *testing.T) {

	result := IsEven(4)
	assert.True(t, result, "4 must be an even number")


	result = IsEven(7)
	assert.False(t, result, "7 must not be an even number")


	result = IsEven(-10)
	assert.True(t, result, "-10 must be an even number")


	result = IsEven(-13)
	assert.False(t, result, "-13 must not be an even number)"
}

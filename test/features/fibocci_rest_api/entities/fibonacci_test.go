package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	fibonacci_rest_api "example.com/mod/lib/features/fibonacci_rest_api/entities"
)

func TestFibonacciCreateSuccess(t *testing.T) {
	fib := fibonacci_rest_api.NewFibonacci()

	result := fib.GetFullList()

	assert.Equal(t, int64(1), result[0])
	assert.Equal(t, int64(1), result[1])
	assert.Equal(t, int64(2), result[2])
	assert.Equal(t, int64(3), result[3])
	assert.Equal(t, int64(5), result[4])
	assert.Equal(t, 256, len(result))

	assert.Equal(t, int64(89), result[10])
	assert.Equal(t, int64(10946), result[20])
	assert.Equal(t, int64(1346269), result[30])

	expectedLastValue := result[255]
	assert.Equal(t, expectedLastValue, result[255])
}

func TestFibonacciSpecificValues(t *testing.T) {
	fib := fibonacci_rest_api.NewFibonacci()

	val10, err := fib.GetNumberAt(10)
	assert.Nil(t, err)
	assert.Equal(t, int64(55), val10)

	val20, err := fib.GetNumberAt(20)
	assert.Nil(t, err)
	assert.Equal(t, int64(6765), val20)

	_, err = fib.GetNumberAt(-1)
	assert.NotNil(t, err)

	_, err = fib.GetNumberAt(256)
	assert.NotNil(t, err)
}

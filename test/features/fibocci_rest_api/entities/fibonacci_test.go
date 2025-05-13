package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	fibonacci_rest_api "example.com/mod/lib/features/fibonacci_rest_api/entities"
)

func TestFibonacciCreateSuccess(t *testing.T) {
	fib := fibonacci_rest_api.NewFibonacci()

	result := fib.GetFullList()

	assert.Equal(t, "1", result[0].String())
	assert.Equal(t, "1", result[1].String())
	assert.Equal(t, "2", result[2].String())
	assert.Equal(t, "3", result[3].String())
	assert.Equal(t, "5", result[4].String())
	assert.Equal(t, 256, len(result))

	assert.Equal(t, "89", result[10].String())
	assert.Equal(t, "10946", result[20].String())
	assert.Equal(t, "1346269", result[30].String())

	expectedLastValue := result[255]
	assert.Equal(t, expectedLastValue.String(), result[255].String())
}

func TestFibonacciSpecificValues(t *testing.T) {
	fib := fibonacci_rest_api.NewFibonacci()

	val10, err := fib.GetNumberAt(10)
	assert.Nil(t, err)
	assert.Equal(t, "55", val10.String())

	val20, err := fib.GetNumberAt(20)
	assert.Nil(t, err)
	assert.Equal(t, "6765", val20.String())

	_, err = fib.GetNumberAt(-1)
	assert.NotNil(t, err)

	_, err = fib.GetNumberAt(256)
	assert.NotNil(t, err)
}

func TestFibonacciLargeValues(t *testing.T) {
	fib := fibonacci_rest_api.NewFibonacci()

	val100, err := fib.GetNumberAt(100)
	assert.Nil(t, err)
	assert.NotNil(t, val100)

	t.Logf("100th Fibonacci number: %s", val100.String())
}

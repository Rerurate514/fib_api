package fibonacci_rest_api

import (
	"fmt"
	"math/big"
)

type Fibonacci struct {
	list []*big.Int
}

func NewFibonacci() *Fibonacci {
	return &Fibonacci{
		list: createFibonacciArray(),
	}
}

func createFibonacciArray() []*big.Int {
	result := make([]*big.Int, 256)

	result[0] = big.NewInt(1)
	result[1] = big.NewInt(1)

	for i := 2; i < len(result); i++ {
		result[i] = new(big.Int).Add(result[i-2], result[i-1])
	}

	return result
}

func (f *Fibonacci) GetNumberAt(index int) (*big.Int, error) {
	if index < 0 || index >= len(f.list) {
		return nil, fmt.Errorf("インデックス %d は範囲外です（0-%d）", index, len(f.list)-1)
	}

	index--

	return new(big.Int).Set(f.list[index]), nil
}

func (f *Fibonacci) GetFullList() []*big.Int {
	listCopy := make([]*big.Int, len(f.list))
	for i, num := range f.list {
		listCopy[i] = new(big.Int).Set(num)
	}
	return listCopy
}

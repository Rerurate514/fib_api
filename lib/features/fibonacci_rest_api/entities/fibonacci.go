package fibonacci_rest_api

import (
	"fmt"
)

type Fibonacci struct {
	list []int64
}

func NewFibonacci() *Fibonacci {
	return &Fibonacci{
		list: createFibonacciArray(),
	}
}

func createFibonacciArray() []int64 {
	result := make([]int64, 256)

	result[0], result[1] = 0, 1

	for i := 2; i < len(result); i++ {
		result[i] = result[i-2] + result[i-1]
	}

	return result
}

func (f *Fibonacci) GetNumberAt(index int) (int64, error) {
	if index < 0 || index >= len(f.list) {
		return 0, fmt.Errorf("インデックス %d は範囲外です（0-%d）", index, len(f.list)-1)
	}

	return f.list[index], nil
}

func (f *Fibonacci) GetFullList() []int64 {
	return append([]int64{}, f.list...)
}

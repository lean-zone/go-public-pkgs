package random

import (
	"math/rand"
)

func Array1dFloat32(d int) []float32 {
	var rl []float32
	for i := 0; i < d; i++ {
		rl = append(rl, rand.Float32())
	}
	return rl
}

func Array2dFloat32(n int, d int) [][]float32 {
	rl := make([][]float32, n)
	for i := 0; i < n; i++ {
		rl[i] = Array1dFloat32(d)
	}
	return rl
}

func Array1dFloat64(d int) []float64 {
	var rl []float64
	for i := 0; i < d; i++ {
		rl = append(rl, rand.Float64())
	}
	return rl
}

func Array2dFloat64(n int, d int) [][]float64 {
	rl := make([][]float64, n)
	for i := 0; i < n; i++ {
		rl[i] = Array1dFloat64(d)
	}
	return rl
}

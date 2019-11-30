package test

import "fmt"

func average(xs []float64) (avg float64) {
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	return
}

func order(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

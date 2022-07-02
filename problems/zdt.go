package problems

import (
	"math"

	"github.com/HgAmiui9/GO-MOEA/numgo"
)

func ZDT1(x_array []float64, var_n float64) (float64, float64) {
	f1 := x_array[0]
	g := 1 + 9*numgo.Sumf64(x_array[1:])/(var_n-1)
	h := 1 - math.Sqrt(f1/g)
	f2 := g * h
	return f1, f2
}

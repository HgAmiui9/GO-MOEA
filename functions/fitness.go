package functions

import (
	"github.com/HgAmiui9/GO-MOEA/problems"
)

func Fitness(pop [][]float64) (f [][]float64) {
	s := len(pop)

	for i := 0; i < s; i++ {
		f = append(f, problems.ZDT1(pop[i], 30))
	}

	return f
}

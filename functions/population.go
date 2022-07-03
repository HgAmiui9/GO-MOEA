package functions

import (
	"math/rand"
	"time"
)

func InitPopulation(size int, length int) [][]float64 {
	var pop [][]float64

	// 随机生成正态分布的初始种群
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < size; i++ {
		var popi []float64
		for j := 0; j < length; j++ {
			// max=1,min=0 -> miu=0.5
			v := r.NormFloat64() + 0.5
			// 越界处理
			if v > 1 {
				v = 1
			}
			if v < 0 {
				v = 0
			}
			popi = append(popi, v)
		}
		pop = append(pop, popi)
	}

	return pop
}

func SinglePointCrossover(pop [][]float64, rate float32) (newPop [][]float64) {
	s, l := len(pop), len(pop[0])

	r := rand.New(rand.NewSource(time.Now().Unix()))
	// use rand sort slice
	sortIndex := r.Perm(s)

	for i := 0; i < s/2; i++ {
		crossIndex := r.Intn(l - 1)
		if r.Float32() < rate {
			for j := crossIndex; j < l; j++ {
				pop[sortIndex[i*2]][j], pop[sortIndex[i*2+1]][j] = pop[sortIndex[i*2+1]][j], pop[sortIndex[i*2]][j]
			}
		}
	}

	newPop = pop
	return newPop
}

// 采用高斯近似变异
// 对于浮点数组成的基因，更适合采用高斯近似变异
func Mutation(pop [][]float64, rate float32) (newPop [][]float64) {
	s, l := len(pop), len(pop[0])

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < s; i++ {
		if r.Float32() < rate {
			mutIndex := r.Intn(l - 1)
			mutValue := r.NormFloat64() + 0.5
			for mutValue < 0 || mutValue > 1 {
				mutValue = r.NormFloat64() + 0.5
			}
			pop[i][mutIndex] = mutValue
		}
	}

	newPop = pop
	return newPop
}

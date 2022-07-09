package functions

import (
	"math"
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

type individual struct {
	Np int         // 被支配的数量
	Sp [][]float64 //支配的个体集合
}

func FastNondominatedSort(pop [][]float64) (newPop [][]float64) {
	fitness := Fitness(pop)

	var F []*individual
	for i, p := range pop {
		// if p>q
		ind := new(individual)
		var sp [][]float64
		np := 0
		for j, q := range pop {
			if i == j {
				continue
			}

			var dominated, disdominated bool
			for d := 0; d < len(fitness[0]); d++ {
				if fitness[i][d] > fitness[j][d] {
					disdominated = true
					break
				}
				if fitness[i][d] < fitness[j][d] {
					dominated = true
				}
			}
			if !disdominated && dominated {
				sp = append(sp, q)
				np++
			}
		}
		ind.Np, ind.Sp = np, sp
		F = append(F, ind)
	}

}

func CrowdingDistance(i int, fitness [][]float64) (newPop [][]float64) {
	for d := 0; d < len(fitness[0]); d++ {
		left, right, min, max := math.MaxFloat64, math.MaxFloat64, 0.0, 0.0
		for x := 0; x < len(fitness); x++ {

			if fitness[x][d] < min {
				min = fitness[x][d]
			}
			if fitness[x][d] > max {
				max = fitness[x][d]
			}

			if x == i {
				continue
			}

			if fitness[i][d]-fitness[x][d] > 0 && fitness[i][d]-fitness[x][d] < left {
				left = fitness[i][d] - fitness[x][d]
			}
			if fitness[x][d]-fitness[i][d] > 0 && fitness[x][d]-fitness[i][d] < right {
				right = fitness[x][d] - fitness[i][d]
			}

		}
	}

}

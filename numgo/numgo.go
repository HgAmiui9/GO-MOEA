package numgo

// 'linspace' will return 'num' evenly distributed samples.
// 'start' The starting value of the sequence.
// 'end' The end value of the sequence, unless endpoint is set to False.
// 'num' Number of samples to generate. Must be non-negative and not null.
// 'endpoint' If True, stop is the last sample. Otherwise, it is not included. Default is True.
func Linspace(start, end float64, num int, endpoint bool) []float64 {
	var res []float64
	if endpoint {
		delta_x := (end - start) / float64(num-1)
		for start <= end {
			res = append(res, start)
			start += delta_x
		}

	} else {
		delta_x := (end - start) / float64(num)
		for start <= end {
			res = append(res, start)
			start += delta_x
		}

	}

	return res
}

// Sum can calcaulate float32/64、int8/16/32/64、slice、array
//TODO: calculate map type
func Sumf64(array []float64) float64 {
	result := 0.0
	for _, a := range array {
		result += a
	}
	return result
}

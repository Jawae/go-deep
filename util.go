package deep

import "math"

func Mean(xx []float64) float64 {
	var sum float64
	for _, x := range xx {
		sum += x
	}
	return sum / float64(len(xx))
}

func Variance(xx []float64) float64 {
	if len(xx) == 1 {
		return 0.0
	}
	m := Mean(xx)

	var variance float64
	for _, x := range xx {
		variance += math.Pow((x - m), 2)
	}

	return variance / float64(len(xx)-1)
}

func StandardDeviation(xx []float64) float64 {
	return math.Sqrt(Variance(xx))
}

// z-score μ=0 σ=1
func Standardize(xx []float64) {
	m := Mean(xx)
	s := StandardDeviation(xx)

	if s == 0 {
		s = 1
	}

	for i, x := range xx {
		xx[i] = (x - m) / s
	}
}

// scale to (0,1)
func Normalize(xx []float64) {
	min, max := Min(xx), Max(xx)
	for i, x := range xx {
		xx[i] = (x - min) / (max - min)
	}
}

func Min(xx []float64) float64 {
	min := xx[0]
	for _, x := range xx {
		if x < min {
			min = x
		}
	}
	return min
}

func Max(xx []float64) float64 {
	max := xx[0]
	for _, x := range xx {
		if x > max {
			max = x
		}
	}
	return max
}

func ArgMax(xx []float64) int {
	max, idx := xx[0], 0
	for i, x := range xx {
		if x > max {
			max, idx = xx[i], i
		}
	}
	return idx
}

func Sgn(x float64) float64 {
	switch {
	case x < 0:
		return -1.0
	case x > 0:
		return 1.0
	}
	return 0
}

func Sum(xx []float64) (sum float64) {
	for _, x := range xx {
		sum += x
	}
	return
}

func Softmax(xx []float64) []float64 {
	out := make([]float64, len(xx))
	var sum float64
	max := Max(xx)
	for i, x := range xx {
		out[i] = math.Exp(x - max)
		sum += out[i]
	}
	for i := range out {
		out[i] /= sum
	}
	return out
}
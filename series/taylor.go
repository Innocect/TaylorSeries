package series

type TaylorSeries interface {
	Calculate(n, x uint64, ch chan uint64) float64
}

type taylorSeries struct{}

func InitTaylorSeries() TaylorSeries {
	return &taylorSeries{}
}

func (t *taylorSeries) Calculate(n, x uint64, ch chan uint64) float64 {

	if x == 0 {
		return float64(1)
	}

	return t.Calculate(n, x-1, nil) + powerOfX(n, x)/factorial(x)
}

func factorial(num uint64) float64 {
	if num == 1 || num == 0 {
		return 1
	}

	return float64(num) * factorial(num-1)
}

func powerOfX(num uint64, power uint64) float64 {
	if power == 1 {
		return float64(num)
	}

	return float64(num) * powerOfX(num, power-1)
}

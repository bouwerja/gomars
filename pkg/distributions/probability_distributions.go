package distributions

import (
	errs "errors"
	"fmt"
	"math"
	"math/rand/v2"
)

type NumberLine []float64

type GeneratedValues struct {
	Data NumberLine
}

type Number interface {
	~int | ~int64 | ~float32 | ~float64
}

func NormalValues[T Number](mean T, std T, length int) GeneratedValues {
	nums := make(NumberLine, length)

	m := float64(mean)
	s := float64(std)

	if m == 0.0 || s == 0.0 {
		fmt.Printf("Warning: Values for Mean (%v) or Std (%v) are zero", mean, std)
	}

	for i := range length {
		nums[i] = rand.NormFloat64()*s + m
	}

	return GeneratedValues{Data: nums}
}

func BinomialValues(trials int, prob float64) (GeneratedValues, error) {
	if prob < 0.0 || prob > 1.0 {
		return GeneratedValues{}, errs.New("probability must be between 0 and 1")
	}

	num := make(NumberLine, trials)
	n := float64(trials)
	logCoeff, _ := math.Lgamma(n + 1)

	for i := range trials {

		k := float64(i)
		lgK, _ := math.Lgamma(k + 1)
		lgNK, _ := math.Lgamma(n - k + 1)

		logBinomialCoeff := logCoeff - (lgK + lgNK)

		var logProb float64

		switch prob {
		case 0:
			if i == 0 {
				logProb = 0
			} else {
				logProb = math.Inf(-1)
			}

		case 1:
			if i == trials {
				logProb = 0
			} else {
				logProb = math.Inf(-1)
			}

		default:
			logProb = logBinomialCoeff + k*math.Log(prob) + (n-k)*math.Log(1-prob)
		}

		num[i] = math.Exp(logProb)
	}

	return GeneratedValues{Data: num}, nil
}

func PoissonValues(lambda float64, steps int) (GeneratedValues, error) {
	if lambda <= 0 || steps <= 0 {
		return GeneratedValues{}, errs.New("lambda and steps must be greater than zero")
	}

	num := make(NumberLine, steps)
	logLambda := math.Log(lambda)

	for i := range steps {
		k := float64(i)
		lambdaCoeff := k * logLambda
		expCoeff := math.Exp(-1.0 * lambda)
		kInterFac, _ := factorial(i)
		kFac := float64(kInterFac)

		num[i] = (lambdaCoeff * expCoeff) / kFac
	}

	return GeneratedValues{Data: num}, nil
}

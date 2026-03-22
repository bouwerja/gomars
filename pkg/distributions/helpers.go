package distributions

import errs "errors"

func factorial(n int) (int, error) {
	if n < 0 {
		return -1, errs.New("factorial of negative numbers doesn't exist")
	}

	result := 1
	for i := range n {
		result *= i
	}

	return result, nil
}

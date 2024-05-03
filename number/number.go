package number

import (
	"math"
	"math/rand/v2"
)

func Euler(p, q uint64) uint64 {
	return (p - 1) * (q - 1)
}

func GCD(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func IsPrime(N uint64) bool {
	if N < 2 {
		return false
	}

	if N == 2 {
		return true
	}

	if N%2 == 0 {
		return false
	}

	for i := uint64(3); i < uint64(math.Sqrt(float64(N)))+1; i = i + 2 {
		if N%i == 0 {
			return false
		}
	}

	return true
}

func Coprime(N uint64) uint64 {
	var min, max uint64 = 2, N - 2

	for {
		a := rand.Uint64N(max-1) + min
		if GCD(N, a) == 1 {
			return a
		}
	}
}

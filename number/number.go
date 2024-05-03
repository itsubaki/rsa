package number

import (
	"crypto/rand"
	"fmt"
	"math"
	randv2 "math/rand/v2"
)

// Euler returns the Euler's function.
// p and q must be prime numbers.
func Euler(p, q int) int {
	return (p - 1) * (q - 1)
}

// GCD returns the greatest common divisor of a and b.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// IsPrime returns true if N is prime number.
func IsPrime(N int) bool {
	if N < 2 {
		return false
	}

	if N == 2 {
		return true
	}

	if N%2 == 0 {
		return false
	}

	for i := 3; i < int(math.Sqrt(float64(N)))+1; i = i + 2 {
		if N%i == 0 {
			return false
		}
	}

	return true
}

// Coprime returns a random coprime number in [2, N).
func Coprime(N int, seed ...[32]byte) (int, error) {
	s, err := NewSource(seed...)
	if err != nil {
		return 0, fmt.Errorf("new source: %v", err)
	}
	rng := randv2.New(s)

	var min, max int = 2, N - 2
	for {
		a := rng.IntN(max-1) + min
		if GCD(N, a) == 1 {
			return a, nil
		}
	}
}

func NewSource(seed ...[32]byte) (randv2.Source, error) {
	if len(seed) == 0 {
		read, err := Read()
		if err != nil {
			return nil, fmt.Errorf("read: %v", err)
		}

		seed = append(seed, read)
	}

	return randv2.NewChaCha8(seed[0]), nil
}

func Read() ([32]byte, error) {
	var p [32]byte
	if _, err := rand.Read(p[:]); err != nil {
		return [32]byte{}, fmt.Errorf("read: %v", err)
	}

	return p, nil
}

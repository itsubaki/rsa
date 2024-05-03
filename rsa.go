package rsa

import "github.com/itsubaki/rsa/number"

func PublicKey(e int, seed ...[32]byte) int {
	return Must(number.Coprime(e, seed...))
}

func PrivateKey(e, E int) int {
	for i := 1; ; i++ {
		if (E*i)%e == 1 {
			return i
		}
	}
}

func Encrypt(msg, E, N int) int {
	return mod(msg, E, N)
}

func Decrypt(msg, D, N int) int {
	return mod(msg, D, N)
}

func mod(msg int, e int, n int) int {
	out := msg
	for i := 1; i < e; i++ {
		out = (out * msg) % n
	}

	return out
}

func Must[T any](a T, err error) T {
	if err != nil {
		panic(err)
	}

	return a
}

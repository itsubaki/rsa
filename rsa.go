package rsa

import "github.com/itsubaki/rsa/number"

// PublicKey returns a public key.
func PublicKey(e int, seed ...[32]byte) int {
	return Must(number.Coprime(e, seed...))
}

// PrivateKey returns a private key.
func PrivateKey(e, E int) int {
	for x := 1; ; x++ {
		if (E*x)%e == 1 {
			return x
		}
	}
}

// Encrypt returns an encrypted message.
func Encrypt(msg, E, N int) int {
	return mod(msg, E, N)
}

// Decrypt returns a decrypted message.
func Decrypt(msg, D, N int) int {
	return mod(msg, D, N)
}

func mod(msg int, e int, n int) int {
	x := msg
	for i := 1; i < e; i++ {
		x = (x * msg) % n
	}

	return x
}

// Must panics if err is not nil.
func Must[T any](a T, err error) T {
	if err != nil {
		panic(err)
	}

	return a
}

# rsa

[![PkgGoDev](https://pkg.go.dev/badge/github.com/itsubaki/rsa)](https://pkg.go.dev/github.com/itsubaki/rsa)
[![Go Report Card](https://goreportcard.com/badge/github.com/itsubaki/rsa?style=flat-square)](https://goreportcard.com/report/github.com/itsubaki/rsa)
[![tests](https://github.com/itsubaki/rsa/workflows/tests/badge.svg?branch=main)](https://github.com/itsubaki/rsa/actions)

## Example

```go
p, q := 17, 19

e := number.Euler(p, q)
E := rsa.PublicKey(e)
D := rsa.PrivateKey(e, E)
N := p * q
fmt.Printf("p=%v, q=%v, e=%v, E=%v, D=%v, N=%v\n", p, q, e, E, D, N)

for msg := 2; msg < N-1; msg++ {
	enc := rsa.Encrypt(msg, E, N)
	dec := rsa.Decrypt(enc, D, N)
	fmt.Printf("msg=%3d, enc=%3d, dec=%3d\n", msg, enc, dec)
}

// p=17, q=19, e=288, E=35, D=107, N=323
// msg=  2, enc=314, dec=  2
// msg=  3, enc=146, dec=  3
// msg=  4, enc= 81, dec=  4
// msg=  5, enc= 23, dec=  5
// msg=  6, enc=301, dec=  6
// msg=  7, enc=258, dec=  7
// msg=  8, enc=240, dec=  8
// msg=  9, enc=321, dec=  9
// msg= 10, enc=116, dec= 10
// msg= 11, enc=311, dec= 11
// msg= 12, enc=198, dec= 12
// ...
```

package number_test

import (
	"fmt"
	"testing"

	"github.com/itsubaki/rsa/number"
)

func ExampleEuler() {
	e := number.Euler(17, 19)
	fmt.Println(e)

	// Output:
	// 288
}

func ExampleGCD() {
	gcd := number.GCD(15, 7)
	fmt.Println(gcd)

	// Output:
	// 1
}

func TestGCD(t *testing.T) {
	cases := []struct {
		a, b uint64
		want uint64
	}{
		{15, 2, 1},
		{15, 4, 1},
		{15, 7, 1},
		{15, 11, 1},
		{15, 13, 1},
		{15, 14, 1},
	}

	for _, c := range cases {
		got := number.GCD(c.a, c.b)
		if got != c.want {
			t.Errorf("got=%v want=%d", got, c.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	cases := []struct {
		in   uint64
		want bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{14, false},
		{15, false},
		{16, false},
		{17, true},
		{18, false},
		{19, true},
		{20, false},
		{21, false},
	}

	for _, c := range cases {
		got := number.IsPrime(c.in)
		if got != c.want {
			t.Errorf("got=%v, want=%v", got, c.want)
		}
	}
}

func ExampleCoprime() {
	p := number.Coprime(15)

	for _, e := range []uint64{2, 4, 7, 8, 11, 13, 14} {
		if p == e {
			fmt.Println("found")
			break
		}
	}

	// Output:
	// found
}

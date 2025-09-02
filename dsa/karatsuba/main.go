package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	n1 := new(big.Int)
	n2 := new(big.Int)
	n1.SetString("3141592653589793238462643383279502884197169399375105820974944592", 10)
	n2.SetString("2718281828459045235360287471352662497757247093699959574966967627", 10)
	expected := new(big.Int)
	expected.Mul(n1, n2)
	fmt.Println("Expected: ", expected)

	karatsubaRes := karatsuba(n1, n2)
	fmt.Println("Karatsuba: ", karatsubaRes)

	fmt.Println("Correct: ", expected.Cmp(karatsubaRes) == 0)
}

func karatsuba(a, b *big.Int) *big.Int {
	// Handle zero quickly
	if a.Sign() == 0 || b.Sign() == 0 {
		return new(big.Int)
	}

	sign := a.Sign() * b.Sign()
	aa := new(big.Int).Abs(a)
	bb := new(big.Int).Abs(b)

	res := karatsubaRec(aa, bb)
	if sign < 0 {
		res.Neg(res)
	}
	return res
}

func karatsubaRec(a, b *big.Int) *big.Int {
	n := int(math.Max(float64(ndigits(a)), float64(ndigits(b))))
	if n <= 2 {
		return new(big.Int).Mul(a, b)
	}
	m := (n + 1) / 2

	// pow10m = 10^m
	pow10m := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(m)), nil)

	// Split: x = x1*10^m + x0
	a1 := new(big.Int).Quo(a, pow10m)
	a0 := new(big.Int).Mod(a, pow10m)
	b1 := new(big.Int).Quo(b, pow10m)
	b0 := new(big.Int).Mod(b, pow10m)

	// z0 = a0*b0
	z0 := karatsubaRec(a0, b0)
	// z2 = a1*b1
	z2 := karatsubaRec(a1, b1)

	// z1 = (a0+a1)*(b0+b1) - z2 - z0
	t1 := new(big.Int).Add(a0, a1)
	t2 := new(big.Int).Add(b0, b1)
	z1 := karatsubaRec(t1, t2)
	z1.Sub(z1, z2)
	z1.Sub(z1, z0)

	// result = z2*10^(2m) + z1*10^m + z0
	pow10_2m := new(big.Int).Mul(pow10m, pow10m)

	res := new(big.Int).Mul(z2, pow10_2m)
	res.Add(res, new(big.Int).Mul(z1, pow10m))
	res.Add(res, z0)
	return res
}

func ndigits(x *big.Int) int {
	return len(x.String())
}

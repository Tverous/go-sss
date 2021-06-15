package sss

import (
	"fmt"
	"math"
	"math/rand"
)

func ReconstructSecret(shares []int) {

}

func MakePolynomial(x int, coefficients []int) int {
	point := 0
	for i, value := range coefficients {
		point += int(math.Pow(float64(x), float64(i))) * value
	}

	return point
}

func Coeff(t int, secret int) []int {
	coff := rand.Perm(t)
	coff = append(coff, secret)

	return coff
}

func GenerateShares(n int, m int, secret int) {
	coefficients := Coeff(m, secret)

	for i := 0; i < n; i ++ {
		x := rand.Intn(30)
		fmt.Printf("x = %d, shares = %d\n", x, MakePolynomial(x, coefficients))
	}
}
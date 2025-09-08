package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := new(big.Int)
	b := new(big.Int)

	a.SetString("123456789012345678901234567890", 10)
	b.SetString("987654321098765432109876543210", 10)

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	mul := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)

	fmt.Println("Сложение =", sum)
	fmt.Println("Вычитание =", diff)
	fmt.Println("Умножение =", mul)
	fmt.Println("Деление =", div)
}

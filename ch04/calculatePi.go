package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
)

var precision uint = 0

func Pi(accuracy uint) *big.Float {
	k := 0
	pi := new(big.Float).SetPrec(precision).SetFloat64(0)
	k1k2k3 := new(big.Float).SetPrec(precision).SetFloat64(0)
	k4k5k6 := new(big.Float).SetPrec(precision).SetFloat64(0)
	temp := new(big.Float).SetPrec(precision).SetFloat64(0)
	minusOne := new(big.Float).SetPrec(precision).SetFloat64(-1)
	total := new(big.Float).SetPrec(precision).SetFloat64(0)

	two2Six := math.Pow(2, 6)
	two2SixBig := new(big.Float).SetPrec(precision).SetFloat64(two2Six)

	for {
		if k > int(accuracy) {
			break
		}
		k1 := new(big.Float).SetPrec(precision)
		k1.Quo(big.NewFloat(1), big.NewFloat(float64(10*k+9)))
		k2 := new(big.Float).SetPrec(precision)
		k2.Quo(big.NewFloat(64), big.NewFloat(float64(10*k+3)))
		k3 := new(big.Float).SetPrec(precision)
		k3.Quo(big.NewFloat(32), big.NewFloat(float64(4*k+1)))
		k1k2k3.Sub(k1, k2)
		k1k2k3.Sub(k1k2k3, k3)

		k4 := new(big.Float).SetPrec(precision)
		k4.Quo(big.NewFloat(4), big.NewFloat(float64(10*k+5)))
		k5 := new(big.Float).SetPrec(precision)
		k5.Quo(big.NewFloat(4), big.NewFloat(float64(10*k+7)))
		k6 := new(big.Float).SetPrec(precision)
		k6.Quo(big.NewFloat(1), big.NewFloat(float64(4*k+3)))
		k4k5k6.Add(k4, k5)
		k4k5k6.Add(k4k5k6, k6)
		k4k5k6 = k4k5k6.Mul(k4k5k6, minusOne)
		temp.Add(k1k2k3, k4k5k6)

		k7temp := new(big.Int).Exp(big.NewInt(-1), big.NewInt(int64(k)), nil)
		k8temp := new(big.Int).Exp(big.NewInt(1024), big.NewInt(int64(k)), nil)

		k7 := new(big.Float).SetPrec(precision).SetFloat64(0)
		k7.SetInt(k7temp)
		k8 := new(big.Float).SetPrec(precision).SetFloat64(0)
		k8.SetInt(k8temp)

		k9 := new(big.Float).SetPrec(precision)
		k9.Quo(big.NewFloat(256), big.NewFloat(float64(10*k+1)))
		k9.Add(k9, temp)
		total.Mul(k9, k7)
		total.Quo(total, k8)
		pi.Add(pi, total)

		k = k + 1
	}
	pi.Quo(pi, two2SixBig)
	return pi
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide one numeric argument!")
		os.Exit(1)
	}

	temp, _ := strconv.ParseUint(arguments[1], 10, 32)
	precision = uint(temp) * 3

	PI := Pi(precision)
	fmt.Println(PI)
}

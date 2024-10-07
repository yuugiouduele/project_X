package main

import (
	"fmt"
	"go/constant"
	"go/token"
	"math"
	"time"
)

func main() {
	zero := constant.MakeInt64(0)
	one := constant.MakeInt64(1)
	negOne := constant.MakeInt64(-1)
	mkComplex := func(a, b constant.Value) constant.Value {
		b = constant.MakeImag(b)
		return constant.BinaryOp(a, token.ADD, b)
	}
	vs := []constant.Value{
		negOne,
		mkComplex(zero, negOne),
		mkComplex(one, negOne),
		mkComplex(negOne, one),
		mkComplex(negOne, negOne),
		zero,
		mkComplex(zero, zero),
		one,
		mkComplex(zero, one),
		mkComplex(one, one),
	}
	for _, v := range vs {
		fmt.Printf("% d %s\n", constant.Sign(v), v)
	}
	//区切り
	x := math.Abs(-2)
	fmt.Printf("%.1f\n", x)
	y := math.Abs(2)
	fmt.Printf("%.1f\n", y)
	z := math.Dim(7, -88)
	fmt.Printf("%.2f\n\n", z)
	j := math.J0(56)
	fmt.Printf("%.2f\n\n", j)
	yn := math.Yn(67, 0.98)
	fmt.Printf("%.2f\n\n", yn)
	po := math.Pow(17, 4)
	fmt.Printf("%.2f\n\n", po)
	t := time.Hour.Microseconds()
	t2 := time.TimeOnly
	t3 := time.Now()
	t4 := time.January
	fmt.Println(t, t2, t3, t4)
}

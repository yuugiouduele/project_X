package main

import (
	"fmt"
)

func crypto(x float64) {
	LogN := 12
	Q := []uint64{0x800004001, 0x40002001}
	P := []uint64{0x4000026001}
	fmt.Println(x, LogN, Q, P)
}

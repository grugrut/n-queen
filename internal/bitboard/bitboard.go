package bitboard

import (
	"fmt"
	"math"
)

var result int

func Solve(N int) {
	if N > 8 {
		fmt.Println("Now version, specify N <= 8")
		return
	}
	var checked uint64
	for i := 0; i < N; i++ {
		checked = (checked << 8) | (uint64(math.Pow(2, float64(8-N)) - 1))
	}
	for i := N; i < 8; i++ {
		checked = (checked << 8) | 0xff
	}

	backtrack(N, 0, checked)
	fmt.Println(result)
}

func check() bool {
	return true
}

func backtrack(N int, q uint64, checked uint64) {
	for {
		b := (getPlacableCell(q) & ^checked)
		if b == 0 {
			if countBit(q) == N {
				fmt.Printf("%064b OK\n", q)
				result++
			}
			return
		}
		next := (b & -b)
		checked |= next
		backtrack(N, q|next, checked)
	}
}

func countBit(i uint64) int {
	i -= ((i >> 1) & 0x5555555555555555)
	i = (i & 0x3333333333333333) + ((i >> 2) & 0x3333333333333333)
	i = (i + (i >> 4)) & 0x0f0f0f0f0f0f0f0f
	i += i >> 8
	i += i >> 16
	i += i >> 32
	return int(i & 0x7f)
}

func getPlacableCell(q uint64) uint64 {
	var b uint64 = 0

	var temp uint64
	var sheild uint64

	// left
	sheild = 0x7f7f7f7f7f7f7f7f
	temp = q
	temp |= (temp & sheild) << 1
	temp |= (temp & sheild) << 1
	temp |= (temp & sheild) << 1
	temp |= (temp & sheild) << 1
	temp |= (temp & sheild) << 1
	temp |= (temp & sheild) << 1
	temp |= (temp & sheild) << 1
	b |= temp

	// up left
	sheild = 0x007f7f7f7f7f7f7f
	temp = q
	temp |= (temp & sheild) << 9
	temp |= (temp & sheild) << 9
	temp |= (temp & sheild) << 9
	temp |= (temp & sheild) << 9
	temp |= (temp & sheild) << 9
	temp |= (temp & sheild) << 9
	temp |= (temp & sheild) << 9
	b |= temp

	// up
	sheild = 0x00ffffffffffffff
	temp = q
	temp |= (temp & sheild) << 8
	temp |= (temp & sheild) << 8
	temp |= (temp & sheild) << 8
	temp |= (temp & sheild) << 8
	temp |= (temp & sheild) << 8
	temp |= (temp & sheild) << 8
	temp |= (temp & sheild) << 8
	b |= temp

	// up right
	sheild = 0x00fefefefefefefe
	temp = q
	temp |= (temp & sheild) << 7
	temp |= (temp & sheild) << 7
	temp |= (temp & sheild) << 7
	temp |= (temp & sheild) << 7
	temp |= (temp & sheild) << 7
	temp |= (temp & sheild) << 7
	temp |= (temp & sheild) << 7
	b |= temp

	b = ^b
	return b
}
